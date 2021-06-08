package app

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	database "github.com/nikola43/ecoapigorm/database"
	middlewares "github.com/nikola43/ecoapigorm/middleware"
	"github.com/nikola43/ecoapigorm/routes"
	"github.com/nikola43/ecoapigorm/socketinstance"
	"github.com/nikola43/ecoapigorm/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"math/rand"
	"time"
)

// MessageObject Basic chat message object
type MessageObject struct {
	Data string `json:"data"`
	From string `json:"from"`
	Room string `json:"room"`
	To   string `json:"to"`
}

// Room Chat Room message object
type Room struct {
	Name  string   `json:"name"`
	UUID  string   `json:"uuid"`
	Users []string `json:"users"`
}

var httpServer *fiber.App
var clients map[string]string
var rooms map[string]*Room


type App struct {
}

func (a *App) Initialize(port string) {

	InitializeDatabase(
		utils.GetEnvVariable("MYSQL_USER"),
		utils.GetEnvVariable("MYSQL_PASSWORD"),
		utils.GetEnvVariable("MYSQL_DATABASE"))

	//database.Migrate()
	//fakedatabase.CreateFakeData()

	fmt.Println(utils.GetEnvVariable("AWS_ACCESS_KEY"))
	fmt.Println(utils.GetEnvVariable("AWS_SECRET_KEY"))
	fmt.Println(utils.GetEnvVariable("AWS_ENDPOINT"))
	fmt.Println(utils.GetEnvVariable("AWS_BUCKET_NAME"))
	fmt.Println(utils.GetEnvVariable("AWS_BUCKET_REGION"))
	fmt.Println(utils.GetEnvVariable("MYSQL_USER"))
	fmt.Println(utils.GetEnvVariable("MYSQL_PASSWORD"))
	fmt.Println(utils.GetEnvVariable("MYSQL_DATABASE"))
	fmt.Println(utils.GetEnvVariable("X_API_KEY"))
	fmt.Println(utils.GetEnvVariable("FROM_EMAIL"))
	fmt.Println(utils.GetEnvVariable("FROM_EMAIL_PASSWORD"))

	InitializeHttpServer(port)
}

func HandleRoutes(api fiber.Router) {
	//app.Use(middleware.Logger())

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
	routes.MultimediaRoutes(api)
	routes.PaymentRoutes(api)
}

func InitializeHttpServer(port string) {
	httpServer = fiber.New(fiber.Config{
		BodyLimit: 2000 * 1024 * 1024, // this is the default limit of 4MB
	})
	httpServer.Use(middlewares.XApiKeyMiddleware)
	httpServer.Use(cors.New(cors.Config{}))

	ws := httpServer.Group("/ws")          // /api/v1

	// Setup the middleware to retrieve the data sent in first GET request
	ws.Use(func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})




	// Pull out in another function
	// all the ikisocket callbacks and listeners
	setupSocketListeners()

	ws.Get("/:id", ikisocket.New(func(kws *ikisocket.Websocket) {
		socketinstance.SocketInstance = kws

		// Retrieve the user id from endpoint
		userId := kws.Params("id")

		// Add the connection to the list of the connected clients
		// The UUID is generated randomly and is the key that allow
		// ikisocket to manage Emit/EmitTo/Broadcast
		clients[userId] = kws.UUID

		fmt.Println(clients[userId])

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

// random room id generator
func generateRoomId() string {
	length := 100
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}

	return string(b)
}

// for beautify purposes we clean the rooms object to a list of rooms
func beautifyRoomsObject(rooms map[string]*Room) []*Room {
	var result []*Room

	for _, room := range rooms {
		result = append(result, room)
	}

	return result
}

// Setup all the ikisocket listeners
// pulled out main function
func setupSocketListeners() {

	// The key for the map is message.to
	clients = make(map[string]string)

	// Rooms will be kept in memory as map
	// for faster and easier handling
	rooms = make(map[string]*Room)

	// Multiple event handling supported
	ikisocket.On(ikisocket.EventConnect, func(ep *ikisocket.EventPayload) {
		fmt.Println(fmt.Sprintf("Connection event 1 - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	// On message event
	ikisocket.On(ikisocket.EventMessage, func(ep *ikisocket.EventPayload) {

		fmt.Println(fmt.Sprintf("Message event - User: %s - Message: %s", ep.Kws.GetStringAttribute("user_id"), string(ep.Data)))

		message := MessageObject{}

		// Unmarshal the json message
		// {
		//  "from": "<user-id>",
		//  "to": "<recipient-user-id>",
		//  "room": "<room-id>",
		//  "data": "hello"
		//}
		err := json.Unmarshal(ep.Data, &message)
		if err != nil {
			fmt.Println(err)
			return
		}

		// If the user is trying to send message
		// into a specific group, iterate over the
		// group user socket UUIDs
		if message.Room != "" {
			// Emit the message to all the room participants
			// iterating on all the uuids
			for _, userId := range rooms[message.Room].Users {
				_ = ep.Kws.EmitTo(clients[userId], ep.Data)
			}

			// Other way can be used EmitToList method
			// if you have a []string of ikisocket uuids
			//
			// ep.Kws.EmitToList(list, data)
			//
			return
		}

		// Emit the message directly to specified user
		err = ep.Kws.EmitTo(clients[message.To], ep.Data)
		if err != nil {
			fmt.Println(err)
		}
	})

	// On disconnect event
	ikisocket.On(ikisocket.EventDisconnect, func(ep *ikisocket.EventPayload) {
		// Remove the user from the local clients
		delete(clients, ep.Kws.GetStringAttribute("user_id"))
		fmt.Println(fmt.Sprintf("Disconnection event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	// On close event
	// This event is called when the server disconnects the user actively with .Close() method
	ikisocket.On(ikisocket.EventClose, func(ep *ikisocket.EventPayload) {
		// Remove the user from the local clients
		delete(clients, ep.Kws.GetStringAttribute("user_id"))
		fmt.Println(fmt.Sprintf("Close event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})

	// On error event
	ikisocket.On(ikisocket.EventError, func(ep *ikisocket.EventPayload) {
		fmt.Println(fmt.Sprintf("Error event - User: %s", ep.Kws.GetStringAttribute("user_id")))
	})
}
