package app

import (
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	middlewares "github.com/nikola43/ecoapigorm/middleware"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/routes"
	"github.com/nikola43/ecoapigorm/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var S3Session *s3.S3
var awsBucketName string
var httpServer *fiber.App

type App struct {
}

func (a *App) Initialize(port string) {
	httpServer := fiber.New()
	api := httpServer.Group("/api") // /api
	v1 := api.Group("/v1")          // /api/v1
	api.Use(middlewares.ApiKeyMiddleware)

	Initializedatabase(
		utils.GetEnvVariable("MYSQL_USER"),
		utils.GetEnvVariable("MYSQL_PASSWORD"),
		utils.GetEnvVariable("MYSQL_DATABASE"))

	InitializeAWSConnection(
		utils.GetEnvVariable("AWS_ACCESS_KEY"),
		utils.GetEnvVariable("AWS_SECRET_KEY"),
		utils.GetEnvVariable("AWS_ENDPOINT"),
		utils.GetEnvVariable("AWS_BUCKET_NAME"),
		utils.GetEnvVariable("AWS_BUCKET_REGION"))

	MigrateDB()

	CreateFakeData()

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
}

func Initializedatabase(user, password, database_name string) {
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

	database.GormDB, err = gorm.Open(mysql.New(mysql.Config{Conn: DB}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func MigrateDB() {
	// DROP
	database.GormDB.Migrator().DropTable(&models.Client{})
	database.GormDB.Migrator().DropTable(&models.Employee{})
	database.GormDB.Migrator().DropTable(&models.Clinic{})
	database.GormDB.Migrator().DropTable(&models.Video{})
	database.GormDB.Migrator().DropTable(&models.Image{})
	database.GormDB.Migrator().DropTable(&models.Heartbeat{})
	database.GormDB.Migrator().DropTable(&models.Streaming{})
	database.GormDB.Migrator().DropTable(&models.Recovery{})
	database.GormDB.Migrator().DropTable(&models.PushNotificationData{})
	database.GormDB.Migrator().DropTable(&models.Promo{})
	database.GormDB.Migrator().DropTable(&models.BankAccount{})
	database.GormDB.Migrator().DropTable(&models.CreditCard{})
	database.GormDB.Migrator().DropTable(&models.PaymentMethod{})

	// CREATE
	database.GormDB.AutoMigrate(&models.Client{})
	database.GormDB.AutoMigrate(&models.Employee{})
	database.GormDB.AutoMigrate(&models.Clinic{})
	database.GormDB.AutoMigrate(&models.Video{})
	database.GormDB.AutoMigrate(&models.Image{})
	database.GormDB.AutoMigrate(&models.Heartbeat{})
	database.GormDB.AutoMigrate(&models.Streaming{})
	database.GormDB.AutoMigrate(&models.Recovery{})
	database.GormDB.AutoMigrate(&models.PushNotificationData{})
	database.GormDB.AutoMigrate(&models.Promo{})
	database.GormDB.AutoMigrate(&models.BankAccount{})
	database.GormDB.AutoMigrate(&models.CreditCard{})
	database.GormDB.AutoMigrate(&models.PaymentMethod{})
}

func CreateFakeData() {
	user := models.Client{Name: "Paulo", Email: "pauloxti@gmail.com", Password: utils.HashAndSalt([]byte("1111111111"))}
	database.GormDB.Create(&user)
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
