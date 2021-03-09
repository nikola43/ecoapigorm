package controllers

import (
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/kicks"
	"github.com/nikola43/ecoapigorm/models/promos"
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
	GormDB.Migrator().DropTable(&promos.Promo{})
	GormDB.Migrator().DropTable(&models.BankAccount{})
	GormDB.Migrator().DropTable(&models.CreditCard{})
	GormDB.Migrator().DropTable(&models.PaymentMethod{})
	GormDB.Migrator().DropTable(&models.CalculatorDetail{})
	GormDB.Migrator().DropTable(&models.Calculator{})
	GormDB.Migrator().DropTable(&kicks.Kick{})
	GormDB.Migrator().DropTable(&models.Company{})

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
	GormDB.AutoMigrate(&promos.Promo{})
	GormDB.AutoMigrate(&models.BankAccount{})
	GormDB.AutoMigrate(&models.CreditCard{})
	GormDB.AutoMigrate(&models.PaymentMethod{})
	GormDB.AutoMigrate(&models.Calculator{})
	GormDB.AutoMigrate(&models.CalculatorDetail{})
	GormDB.AutoMigrate(&kicks.Kick{})
	GormDB.AutoMigrate(&models.Company{})
}

func CreateFakeData() {


	// EMPLOYEES ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	employee1 := models.Employee{Name: "Paulo", LastName: "Soares", Phone: "666666666", Email: "pauloxti@gmail.com", Password: utils.HashPassword([]byte("paulo")), Role: "admin"}
	GormDB.Create(&employee1)

	employee2 := models.Employee{Name: "Migue", LastName: "Barrera", Phone: "999999999", Email: "migue@gmail.com", Password: utils.HashPassword([]byte("migue")), Role: "employeee", ParentEmployeeID: employee1.ID}
	GormDB.Create(&employee2)

	employee3 := models.Employee{Name: "Pablo", LastName: "Gutierrez", Phone: "777777777", Email: "pablojoseguit@gmail.com", Password: utils.HashPassword([]byte("pablo")), Role: "employee", ParentEmployeeID: employee1.ID}
	GormDB.Create(&employee3)

	// COMPANIES ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	company1 := models.Company{Name: "Paulo Company", EmployeeID: employee1.ID}
	GormDB.Create(&company1)
	GormDB.Model(&employee1).Update("company_id", company1.ID)
	GormDB.Model(&employee2).Update("company_id", company1.ID)
	GormDB.Model(&employee3).Update("company_id", company1.ID)



	// CLINIC ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	clinic1 := models.Clinic{Name: "Paulo Clinic", EmployeeID: employee1.ID}
	GormDB.Create(&clinic1)

	clinic2 := models.Clinic{Name: "Migue Clinic", EmployeeID: employee2.ID}
	GormDB.Create(&clinic2)

	clinic3 := models.Clinic{Name: "Pablo Clinic", EmployeeID: employee3.ID}
	GormDB.Create(&clinic3)

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
	for i := 1; i < 20; i++ {
		images := models.Image{
			ClientID: 2,
			Url:      "https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA VICENTE (10).png",
			Size:     0,
		}

		GormDB.Create(&images)
	}

	for i := 1; i < 20; i++ {
		videos := models.Video{
			ClientID:     client2.ID,
			Url:          "https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA VICENTE (2) (online-video-cutter.com) copy.mp4-audio.mp4",
			ThumbnailUrl: "https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA VICENTE (2) (online-video-cutter.com) copy.mp4-audio.mp4-thumbnail.jpg",
			Size:         0,
		}

		GormDB.Create(&videos)
	}

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
	
	//PROMOS
	for i := 1; i < 100; i++ {
		promo := promos.Promo{
			ClinicID: clinic2.ID,
			Title:    "Tu primera eco gratis",
			Text:     "Ven a vernos y consigue que te hagamos la primera eco gratis.",
			ImageUrl: "https://s3.eu-central-1.wasabisys.com/stela/weeks/21SM.jpg",
		}

		GormDB.Create(&promo)
	}
}
