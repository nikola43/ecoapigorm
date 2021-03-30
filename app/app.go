package app

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	database "github.com/nikola43/ecoapigorm/database"
	middlewares "github.com/nikola43/ecoapigorm/middleware"
	"github.com/nikola43/ecoapigorm/routes"
	"github.com/nikola43/ecoapigorm/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var httpServer *fiber.App

type App struct {
}

func (a *App) Initialize(port string) {


	InitializeDatabase(
		utils.GetEnvVariable("MYSQL_USER"),
		utils.GetEnvVariable("MYSQL_PASSWORD"),
		utils.GetEnvVariable("MYSQL_DATABASE"))

	database.Migrate()
	// fakedatabase.CreateFakeData()

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

	database.GormDB, err = gorm.Open(mysql.New(mysql.Config{Conn: DB}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err)
	}
}


