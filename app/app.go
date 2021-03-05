package app

import (
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

var S3Session *s3.S3
var awsBucketName string
var httpServer *fiber.App

type App struct {
}

func (a *App) Initialize(port string) {
	httpServer = fiber.New()
	httpServer.Use(cors.New())

	api := httpServer.Group("/api") // /api
	v1 := api.Group("/v1")          // /api/v1
	api.Use(middlewares.ApiKeyMiddleware)

	InitializeDatabase(
		utils.GetEnvVariable("MYSQL_USER"),
		utils.GetEnvVariable("MYSQL_PASSWORD"),
		utils.GetEnvVariable("MYSQL_DATABASE"))

	InitializeAWSConnection(
		utils.GetEnvVariable("AWS_ACCESS_KEY"),
		utils.GetEnvVariable("AWS_SECRET_KEY"),
		utils.GetEnvVariable("AWS_ENDPOINT"),
		utils.GetEnvVariable("AWS_BUCKET_NAME"),
		utils.GetEnvVariable("AWS_BUCKET_REGION"))

	database.Migrate()

	database.CreateFakeData()

	HandleRoutes(v1)

	httpServer.Listen(port)
}

func HandleRoutes(api fiber.Router) {

	//app.Use(middleware.Logger())

	// use routes
	routes.ClientRoutes(api)
	routes.ClinicRoutes(api)
	routes.AuthRoutes(api)
	routes.SignUpRoutes(api)
	routes.CalculatorRoutes(api)
	routes.KickRoutes(api)
	routes.EmployeeRoutes(api)
	routes.CompanyRoutes(api)
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

func InitializeAWSConnection(access_key, secret_key, endpoint, bucket_name, bucket_region string) {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(access_key, secret_key, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(bucket_region),
		S3ForcePathStyle: aws.Bool(true),
	}
	newSession := session.New(s3Config)
	S3Session = s3.New(newSession)
	awsBucketName = bucket_name
}
