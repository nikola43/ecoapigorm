package app

import (
	"database/sql"
	"fmt"
	"github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	database "github.com/nikola43/ecoapigorm/database"
	middlewares "github.com/nikola43/ecoapigorm/middleware"
	"github.com/nikola43/ecoapigorm/routes"
	"github.com/nikola43/ecoapigorm/utils"
	"github.com/nikola43/ecoapigorm/wasabis3manager"
	"github.com/nikola43/ecoapigorm/websockets"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var httpServer *fiber.App

type App struct {
}

func (a *App) Initialize(port string) {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	PROD := os.Getenv("PROD")

	MYSQL_USER := os.Getenv("MYSQL_USER")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")

	S3_ACCESS_KEY := os.Getenv("S3_ACCESS_KEY")
	S3_SECRET_KEY := os.Getenv("S3_SECRET_KEY")
	S3_ENDPOINT := os.Getenv("S3_ENDPOINT")
	S3_BUCKET_REGION := os.Getenv("S3_BUCKET_REGION")

	XAPIKEY := os.Getenv("XAPIKEY")
	FROM_EMAIL := os.Getenv("FROM_EMAIL")
	FROM_EMAIL_PASSWORD := os.Getenv("FROM_EMAIL_PASSWORD")

	_ = XAPIKEY
	_ = FROM_EMAIL
	_ = FROM_EMAIL_PASSWORD

	if PROD == "1" {
		MYSQL_USER = os.Getenv("MYSQL_USER_DEV")
		MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD_DEV")
		MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE_DEV")

		S3_ACCESS_KEY = os.Getenv("S3_ACCESS_KEY_DEV")
		S3_SECRET_KEY = os.Getenv("S3_SECRET_KEY_DEV")
		S3_ENDPOINT = os.Getenv("S3_ENDPOINT_DEV")
		S3_BUCKET_REGION = os.Getenv("S3_BUCKET_REGION_DEV")

		XAPIKEY = os.Getenv("XAPIKEYDEV")
		FROM_EMAIL = os.Getenv("FROM_EMAIL_DEV")
		FROM_EMAIL_PASSWORD = os.Getenv("FROM_EMAIL_PASSWORD_DEV")
	}

	InitializeDatabase(
		MYSQL_USER,
		MYSQL_PASSWORD,
		MYSQL_DATABASE)

	// database.Migrate()
	//fakedatabase.CreateFakeData()

	wasabis3manager.WasabiS3Client = utils.New(
		S3_ACCESS_KEY,
		S3_SECRET_KEY,
		S3_ENDPOINT,
		S3_BUCKET_REGION)

	InitializeHttpServer(port)
}

func HandleRoutes(api fiber.Router) {
	//api.Use(middleware.Logger())

	routes.ClientRoutes(api)
	routes.ClinicRoutes(api)
	routes.AuthRoutes(api)
	routes.SignUpRoutes(api)
	routes.CalculatorRoutes(api)
	routes.KickRoutes(api)
	routes.EmployeeRoutes(api)
	routes.CompanyRoutes(api)
	routes.PromoRoutes(api)
	routes.StreamingRoutes(api)
	multimedia := api.Group("/multimedia")
	routes.MultimediaClinicRoutes(multimedia)
	routes.MultimediaClientRoutes(api)
	routes.PaymentRoutes(api)
}

func InitializeHttpServer(port string) {
	httpServer = fiber.New(fiber.Config{
		BodyLimit: 2000 * 1024 * 1024, // this is the default limit of 4MB
	})
	/*
	//httpServer.Use(middlewares.XApiKeyMiddleware)
	httpServer.Use(cors.New(cors.Config{
		AllowOrigins: "https://panel.ecox.stelast.com",
	}))
	*/

	httpServer.Use(cors.New(cors.Config{
		AllowOrigins: "https://panel.ecox.stelast.com",
	}))

	ws := httpServer.Group("/ws")

	// Setup the middleware to retrieve the data sent in first GET request
	ws.Use(middlewares.WebSocketUpgradeMiddleware)

	// Pull out in another function
	// all the ikisocket callbacks and listeners
	setupSocketListeners()

	ws.Get("/:id", ikisocket.New(func(kws *ikisocket.Websocket) {
		websockets.SocketInstance = kws

		// Retrieve the user id from endpoint
		userId := kws.Params("id")

		// Add the connection to the list of the connected clients
		// The UUID is generated randomly and is the key that allow
		// ikisocket to manage Emit/EmitTo/Broadcast
		websockets.SocketClients[userId] = kws.UUID

		// Every websocket connection has an optional session key => value storage
		kws.SetAttribute("user_id", userId)

		//Broadcast to all the connected users the newcomer
		// kws.Broadcast([]byte(fmt.Sprintf("New user connected: %s and UUID: %s", userId, kws.UUID)), true)
		//Write welcome message
		kws.Emit([]byte(fmt.Sprintf("Socket connected")))
	}))

	api := httpServer.Group("/api") // /api
	v1 := api.Group("/v1")          // /api/v1
	HandleRoutes(v1)

	err := httpServer.Listen(port)
	if err != nil {
		log.Fatal(err)
	}
}

func InitializeDatabase(user, password, database_name string) {
	connectionString := fmt.Sprintf(
		"%s:%s@/%s?parseTime=true",
		user,
		password,
		database_name,
	)

	DB, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	database.GormDB, err = gorm.Open(mysql.New(mysql.Config{Conn: DB}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Fatal(err)
	}
}

// Setup all the ikisocket listeners
// pulled out main function
func setupSocketListeners() {

	// Multiple event handling supported
	ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
		fmt.Println(fmt.Sprintf("Connection socket event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	// On message event
	ikisocket.On(ikisocket.EventMessage, func(ep *ikisocket.EventPayload) {
		fmt.Println(fmt.Sprintf("Message socket event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	// On disconnect event
	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		// Remove the user from the local clients
		delete(websockets.SocketClients, ep.Kws.GetStringAttribute("user_id"))
		fmt.Println(fmt.Sprintf("Disconnection event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	// On close event
	// This event is called when the server disconnects the user actively with .Close() method
	ikisocket.On(ikisocket.EventClose, func(ep *ikisocket.EventPayload) {
		// Remove the user from the local clients
		delete(websockets.SocketClients, ep.Kws.GetStringAttribute("user_id"))
		fmt.Println(fmt.Sprintf("Close event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	// On error event
	ikisocket.On(ikisocket.EventError, func(ep *ikisocket.EventPayload) {
		fmt.Println(fmt.Sprintf("Error event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})
}
