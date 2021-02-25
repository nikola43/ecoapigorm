package app

import (
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	database "github.com/nikola43/ecoapigorm/database"
	models "github.com/nikola43/ecoapigorm/models/responses"
	"github.com/nikola43/ecoapigorm/routes"
	"gorm.io/gorm"
	"log"
	"gorm.io/driver/mysql"
	"os"
)

var S3Session *s3.S3
var awsBucketName string
var httpServer *fiber.App

type App struct {
}


// use godot package to load/read the .env file and
// return the value of the key
func GetEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func (a *App) Initialize(port string) {
	fmt.Println("INIT")

	/* INITIALIZE HTTP SERVER */
	httpServer := fiber.New()
	api := httpServer.Group("/api") // /api
	v1 := api.Group("/v1")          // /api/v1

	/* INITIALIZE DB CONNECTION */
	connectionString := fmt.Sprintf(
		"%s:%s@/%s",
		GetEnvVariable("MYSQL_USER"),
		GetEnvVariable("MYSQL_PASSWORD"),
		GetEnvVariable("MYSQL_DATABASE"))

	DB, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	database.DB, err = gorm.Open(mysql.New(mysql.Config{Conn: DB}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//app.Use(middleware.Logger())

	database.DB.Migrator().DropTable(&models.Client{})
	database.DB.Migrator().DropTable(&models.Employee{})
	database.DB.Migrator().DropTable(&models.Clinic{})
	database.DB.Migrator().DropTable(&models.Video{})
	database.DB.Migrator().DropTable(&models.Image{})
	database.DB.Migrator().DropTable(&models.Heartbeat{})
	database.DB.Migrator().DropTable(&models.Streaming{})

	database.DB.AutoMigrate(
		&models.Client{},
		&models.Employee{},
		&models.Clinic{},
		&models.Video{},
		&models.Image{},
		&models.Heartbeat{},
		&models.Streaming{})


	/* CREATE S3 SESSION */
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(GetEnvVariable("AWS_ACCESS_KEY"), GetEnvVariable("AWS_SECRET_KEY"), ""),
		Endpoint:         aws.String(GetEnvVariable("AWS_ENDPOINT")),
		Region:            aws.String(GetEnvVariable("AWS_BUCKET_REGION")),
		S3ForcePathStyle: aws.Bool(true),
	}
	fmt.Println(GetEnvVariable("AWS_BUCKET_REGION"))
	newSession := session.New(s3Config)
	S3Session = s3.New(newSession)
	awsBucketName = GetEnvVariable("AWS_BUCKET_NAME")


	/* HANDLE ROUTES */
	httpServer.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// use routes
	routes.ClientRoutes(v1)
	routes.ClinicRoutes(v1)
	routes.AuthRoutes(v1)

	httpServer.Listen(port)
}
