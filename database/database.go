package controllers

import (
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/kicks"
	"github.com/nikola43/ecoapigorm/utils"
	"gorm.io/gorm"
	"math/rand"
	"time"
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
	GormDB.Migrator().DropTable(&models.Calculator{})
	GormDB.Migrator().DropTable(&kicks.Kick{})

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
	GormDB.AutoMigrate(&models.Calculator{})
	GormDB.AutoMigrate(&models.CalculatorDetail{})
	GormDB.AutoMigrate(&kicks.Kick{})
	GormDB.AutoMigrate(&kicks.Kick{})
}

func CreateFakeData() {
	// EMPLOYEES ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	employee1 := models.Employee{Name: "Paulo", LastName: "Soares", Phone: "666666666", Email: "pauloxti@gmail.com", Password: utils.HashPassword([]byte("paulo")), Role: "admin"}
	GormDB.Create(&employee1)

	employee2 := models.Employee{Name: "Migue", LastName: "Barrera", Phone: "999999999", Email: "migue@gmail.com", Password: utils.HashPassword([]byte("migue")), Role: "employeee", ParentEmployeeID: 1}
	GormDB.Create(&employee2)

	// CLINIC ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	clinic1 := models.Clinic{Name: "P Clinic", EmployeeID: employee1.ID}
	GormDB.Create(&clinic1)

	clinic2 := models.Clinic{Name: "M Clinic", EmployeeID: employee2.ID}
	GormDB.Create(&clinic2)

	// CLIENTS ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	client1 := models.Client{ClinicID: clinic1.ID, Name: "Paulo", LastName: "Soares", Phone: "666666666", Email: "pauloxti@gmail.com", Password: utils.HashPassword([]byte("paulo"))}
	GormDB.Create(&client1)

	client2 := models.Client{ClinicID: clinic2.ID, Name: "Migue", LastName: "Barrera", Phone: "999999999", Email: "migue@gmail.com", Password: utils.HashPassword([]byte("migue"))}
	GormDB.Create(&client2)

	calculator1 := models.Calculator{ClientID: client1.ID, Week: 4}
	GormDB.Create(&calculator1)

	calculator2 := models.Calculator{ClientID: client2.ID, Week: 8}
	GormDB.Create(&calculator2)

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

	// KICKS
	Time := time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC)
	for i := 1; i < 8; i++ {
		Time = Time.AddDate(0, 1, 0)
		randomKicksCounter := rand.Intn(100)
		for i := 1; i < randomKicksCounter; i++ {

			kick := kicks.Kick{
				Date:     Time.AddDate(0, 0, 0),
				ClientId: 2,
			}

			GormDB.Create(&kick)
		}
	}
}
