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
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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
	httpServer := fiber.New()
	api := httpServer.Group("/api") // /api
	v1 := api.Group("/v1")          // /api/v1

	InitializeDbCorrection(
		GetEnvVariable("MYSQL_USER"),
		GetEnvVariable("MYSQL_PASSWORD"),
		GetEnvVariable("MYSQL_DATABASE"))

	MigrateDb()

	InitializeAWSConnection(
		GetEnvVariable("AWS_ACCESS_KEY"),
		GetEnvVariable("AWS_SECRET_KEY"),
		GetEnvVariable("AWS_ENDPOINT"),
		GetEnvVariable("AWS_BUCKET_NAME"),
		GetEnvVariable("AWS_BUCKET_REGION"))

	HandleRoutes(v1)

	httpServer.Listen(port)
}

func HandleRoutes(api fiber.Router) {

	//app.Use(middleware.Logger())

	// use routes
	routes.ClientRoutes(api)
	routes.ClinicRoutes(api)
	routes.AuthRoutes(api)
}

func InitializeDbCorrection(user, password, database_name string) {
	connectionString := fmt.Sprintf(
		"%s:%s@/%s",
		user,
		password,
		database_name,
	)

	DB, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	database.DB, err = gorm.Open(mysql.New(mysql.Config{Conn: DB}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func MigrateDb() {
	database.DB.Migrator().DropTable(&models.Client{})
	database.DB.Migrator().DropTable(&models.Employee{})
	database.DB.Migrator().DropTable(&models.Clinic{})
	database.DB.Migrator().DropTable(&models.Video{})
	database.DB.Migrator().DropTable(&models.Image{})
	database.DB.Migrator().DropTable(&models.Heartbeat{})
	database.DB.Migrator().DropTable(&models.Streaming{})
	database.DB.Migrator().DropTable(&models.Recovery{})
	database.DB.Migrator().DropTable(&models.PushNotificationData{})
	database.DB.Migrator().DropTable(&models.Promo{})

	database.DB.AutoMigrate(
		&models.Client{},
		&models.Employee{},
		&models.Clinic{},
		&models.Video{},
		&models.Image{},
		&models.Heartbeat{},
		&models.Streaming{},
		&models.Recovery{},
		&models.PushNotificationData{},
		&models.Promo{})
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
