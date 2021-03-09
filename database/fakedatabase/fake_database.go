package fakedatabase

import (
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/models/kicks"
	"github.com/nikola43/ecoapigorm/models/promos"
	"github.com/nikola43/ecoapigorm/models/streaming"
	"github.com/nikola43/ecoapigorm/utils"
	"math/rand"
	"time"
)

func CreateFakeData() {

	// EMPLOYEES ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	employee1 := models.Employee{Name: "Paulo", LastName: "Soares", Phone: "666666666", Email: "pauloxti@gmail.com", Password: utils.HashPassword([]byte("paulo")), Role: "admin"}
	database.GormDB.Create(&employee1)

	employee2 := models.Employee{Name: "Migue", LastName: "Barrera", Phone: "999999999", Email: "migue@gmail.com", Password: utils.HashPassword([]byte("migue")), Role: "employeee", ParentEmployeeID: employee1.ID}
	database.GormDB.Create(&employee2)

	employee3 := models.Employee{Name: "Pablo", LastName: "Gutierrez", Phone: "777777777", Email: "pablojoseguit@gmail.com", Password: utils.HashPassword([]byte("pablo")), Role: "employee", ParentEmployeeID: employee1.ID}
	database.GormDB.Create(&employee3)

	// COMPANIES ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	company1 := models.Company{Name: "Paulo Company", EmployeeID: employee1.ID}
	database.GormDB.Create(&company1)
	database.GormDB.Model(&employee1).Update("company_id", company1.ID)
	database.GormDB.Model(&employee2).Update("company_id", company1.ID)
	database.GormDB.Model(&employee3).Update("company_id", company1.ID)

	// CLINIC ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	clinic1 := models.Clinic{Name: "Paulo Clinic", EmployeeID: employee1.ID}
	database.GormDB.Create(&clinic1)

	clinic2 := models.Clinic{Name: "Migue Clinic", EmployeeID: employee2.ID}
	database.GormDB.Create(&clinic2)

	clinic3 := models.Clinic{Name: "Pablo Clinic", EmployeeID: employee3.ID}
	database.GormDB.Create(&clinic3)

	// CLIENTS ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	client1 := models.Client{ClinicID: clinic1.ID, Name: "Paulo", LastName: "Soares", Phone: "666666666", Email: "pauloxti@gmail.com", Password: utils.HashPassword([]byte("paulo"))}
	database.GormDB.Create(&client1)

	client2 := models.Client{ClinicID: clinic2.ID, Name: "Migue", LastName: "Barrera", Phone: "999999999", Email: "migue@gmail.com", Password: utils.HashPassword([]byte("migue"))}
	database.GormDB.Create(&client2)

	// CALCULATOR ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	for i := 1; i < 41; i++ {
		calculatorDetail := models.CalculatorDetail{
			Week:  uint(i),
			Image: "https://s3.eu-central-1.wasabisys.com/stela/weeks/21SM.jpg",
			Text:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book",
		}

		database.GormDB.Create(&calculatorDetail)
	}

	CreateFakeByClient(client1.ID)
 	CreateFakeByClient(client2.ID)

	CreateFakeByClinic(clinic2.ID)
	CreateFakeByClinic(clinic2.ID)
}

func CreateFakeByClinic(clinicId uint) {
	//PROMOS
	for i := 1; i < 100; i++ {
		promo := promos.Promo{
			ClinicID: clinicId,
			Title:    "Tu primera eco gratis",
			Text:     "Ven a vernos y consigue que te hagamos la primera eco gratis.",
			ImageUrl: "https://s3.eu-central-1.wasabisys.com/stela/weeks/21SM.jpg",
		}

		database.GormDB.Create(&promo)
	}

}
func CreateFakeByClient(clientId uint) {
	calculator1 := models.Calculator{ClientID: clientId, Week: 4}
	database.GormDB.Create(&calculator1)

	calculator2 := models.Calculator{ClientID: clientId, Week: 8}
	database.GormDB.Create(&calculator2)

	// IMAGES ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	for i := 1; i < 20; i++ {
		images := models.Image{
			ClientID: clientId,
			Url:      "https://s3.eu-central-1.wasabisys.com/babyandme/4/image/MARIANA VICENTE (10).png",
			Size:     0,
		}

		database.GormDB.Create(&images)
	}

	for i := 1; i < 20; i++ {
		videos := models.Video{
			ClientID:     clientId,
			Url:          "https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA VICENTE (2) (online-video-cutter.com) copy.mp4-audio.mp4",
			ThumbnailUrl: "https://s3.eu-central-1.wasabisys.com/babyandme/4/video/MARIANA VICENTE (2) (online-video-cutter.com) copy.mp4-audio.mp4-thumbnail.jpg",
			Size:         0,
		}

		database.GormDB.Create(&videos)
	}

	// KICKS
	Time := time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC)
	for i := 1; i < 8; i++ {
		Time = Time.AddDate(0, 1, 0)
		randomKicksCounter := rand.Intn(100)
		for i := 1; i < randomKicksCounter; i++ {

			kick := kicks.Kick{
				Date:     Time.AddDate(0, 0, 0),
				ClientId: clientId,
			}

			database.GormDB.Create(&kick)
		}
	}

	//STREAMING
	for i := 1; i < 100; i++ {
		streaming := streaming.Streaming{
			ClientID:        clientId,
			Url:             "https://www.youtube.com/watch?v=5qap5aO4i9A",
			Code:            "12345",
		}

		database.GormDB.Create(&streaming)
	}
}
