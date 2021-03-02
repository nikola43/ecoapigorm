package controllers

import (
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func Migrate() {
	// DROP
	GormDB.Migrator().DropTable(&models.Client{})
	GormDB.Migrator().DropTable(&models.Employee{})
	GormDB.Migrator().DropTable(&models.Clinic{})
	GormDB.Migrator().DropTable(&models.Video{})
	GormDB.Migrator().DropTable(&models.Image{})
	GormDB.Migrator().DropTable(&models.Heartbeat{})
	GormDB.Migrator().DropTable(&models.Streaming{})
	GormDB.Migrator().DropTable(&models.Recovery{})
	GormDB.Migrator().DropTable(&models.PushNotificationData{})
	GormDB.Migrator().DropTable(&models.Promo{})
	GormDB.Migrator().DropTable(&models.BankAccount{})
	GormDB.Migrator().DropTable(&models.CreditCard{})
	GormDB.Migrator().DropTable(&models.PaymentMethod{})
	GormDB.Migrator().DropTable(&models.CalculatorDetail{})

	// CREATE
	GormDB.AutoMigrate(&models.Client{})
	GormDB.AutoMigrate(&models.Employee{})
	GormDB.AutoMigrate(&models.Clinic{})
	GormDB.AutoMigrate(&models.Video{})
	GormDB.AutoMigrate(&models.Image{})
	GormDB.AutoMigrate(&models.Heartbeat{})
	GormDB.AutoMigrate(&models.Streaming{})
	GormDB.AutoMigrate(&models.Recovery{})
	GormDB.AutoMigrate(&models.PushNotificationData{})
	GormDB.AutoMigrate(&models.Promo{})
	GormDB.AutoMigrate(&models.BankAccount{})
	GormDB.AutoMigrate(&models.CreditCard{})
	GormDB.AutoMigrate(&models.PaymentMethod{})
	GormDB.AutoMigrate(&models.CalculatorDetail{})
}

func CreateFakeData() {
	// CLIENTS ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	client1 := models.Client{Name: "Paulo", Email: "pauloxti@gmail.com", Password: utils.HashPassword([]byte("paulo"))}
	GormDB.Create(&client1)

	client2 := models.Client{Name: "Migue", Email: "migue@gmail.com", Password: utils.HashPassword([]byte("migue"))}
	GormDB.Create(&client2)

	// EMPLOYEES ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	user1 := models.Employee{Name: "Paulo", Email: "pauloxti@gmail.com", Password: utils.HashPassword([]byte("paulo")), Role: "admin"}
	GormDB.Create(&user1)

	user2 := models.Employee{Name: "Migue", Email: "migue@gmail.com", Password: utils.HashPassword([]byte("migue")), Role: "employeee", ParentEmployeeID: 1}
	GormDB.Create(&user2)

	// IMAGES ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	images := make([]models.Image, 0)
	images = append(images, models.Image{
		ClientID: 2,
		Url:      "https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA VICENTE (10).png",
		Size:     0,
	})

	images = append(images, models.Image{
		ClientID: 2,
		Url:      "https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA VICENTE (13).png",
		Size:     0,
	})

	GormDB.Create(&images)

	videos := make([]models.Video, 0)
	videos = append(videos, models.Video{
		ClientID:     2,
		Url:          "https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA VICENTE (2) (online-video-cutter.com) copy.mp4-audio.mp4",
		ThumbnailUrl: "https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA VICENTE (2) (online-video-cutter.com) copy.mp4-audio.mp4-thumbnail.jpg",
		Size:         0,
	})

	videos = append(videos, models.Video{
		ClientID:     2,
		Url:          "https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA VICENTE (2).mp4-audio copy.mp4-audio.mp4",
		ThumbnailUrl: "https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA VICENTE (2).mp4-audio copy.mp4-audio.mp4-thumbnail.jpg",
		Size:         0,
	})

	GormDB.Create(&videos)

	// CALCULATOR ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	for i := 1; i < 41; i++ {
		calculatorDetail := models.CalculatorDetail{
			Week:  uint(i),
			Image: "https://s3.eu-central-1.wasabisys.com/stela/weeks/21SM.jpg",
			Text:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book",
		}

		GormDB.Create(&calculatorDetail)
	}
}
