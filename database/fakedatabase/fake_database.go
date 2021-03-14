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

	employee3 := models.Employee{Name: "Pablo", LastName: "Gutierrez", Phone: "777777777", Email: "pablojoseguit@gmail.com", Password: utils.HashPassword([]byte("pablo")), Role: "admin", ParentEmployeeID: employee1.ID}
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
	client1 := models.Client{ClinicID: clinic1.ID, Name: "Paulo", LastName: "Soares", Phone: "666 666 666", Email: "paulo@gmail.com", Password: utils.HashPassword([]byte("paulo"))}
	database.GormDB.Create(&client1)

	client2 := models.Client{ClinicID: clinic1.ID, Name: "Migue", LastName: "Barrera", Phone: "999 999 999", Email: "migue@gmail.com", Password: utils.HashPassword([]byte("migue"))}
	database.GormDB.Create(&client2)

	client3 := models.Client{ClinicID: clinic1.ID, Name: "Marta", LastName: "Martín", Phone: "999 999 999", Email: "marta@gmail.com", Password: utils.HashPassword([]byte("marta"))}
	database.GormDB.Create(&client3)

	client4 := models.Client{ClinicID: clinic1.ID, Name: "Fernanda", LastName: "Portal", Phone: "999 999 999", Email: "fernanda@gmail.com", Password: utils.HashPassword([]byte("fernanda"))}
	database.GormDB.Create(&client4)

	client5 := models.Client{ClinicID: clinic1.ID, Name: "Alejandra", LastName: "Fernández", Phone: "999 999 999", Email: "alejandra@gmail.com", Password: utils.HashPassword([]byte("alejandra"))}
	database.GormDB.Create(&client5)

	client6 := models.Client{ClinicID: clinic1.ID, Name: "Claudia", LastName: "Sánchez", Phone: "999 999 999", Email: "claudia@gmail.com", Password: utils.HashPassword([]byte("claudia"))}
	database.GormDB.Create(&client6)

	client7 := models.Client{ClinicID: clinic1.ID, Name: "Sofía", LastName: "González", Phone: "999 999 999", Email: "sofia@gmail.com", Password: utils.HashPassword([]byte("sofia"))}
	database.GormDB.Create(&client7)

	client8 := models.Client{ClinicID: clinic1.ID, Name: "María", LastName: "Ruiz", Phone: "999 999 999", Email: "maria@gmail.com", Password: utils.HashPassword([]byte("maria"))}
	database.GormDB.Create(&client8)

	client9 := models.Client{ClinicID: clinic1.ID, Name: "Ana", LastName: "Diaz", Phone: "999 999 999", Email: "ana@gmail.com", Password: utils.HashPassword([]byte("ana"))}
	database.GormDB.Create(&client9)

	client10 := models.Client{ClinicID: clinic1.ID, Name: "Inma", LastName: "Romero", Phone: "999 999 999", Email: "inma@gmail.com", Password: utils.HashPassword([]byte("inma"))}
	database.GormDB.Create(&client10)

	client11 := models.Client{ClinicID: clinic1.ID, Name: "Mónica", LastName: "Navarro", Phone: "999 999 999", Email: "monica@gmail.com", Password: utils.HashPassword([]byte("monica"))}
	database.GormDB.Create(&client11)

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
	CreateFakeByClient(client3.ID)
	CreateFakeByClient(client4.ID)
	CreateFakeByClient(client5.ID)
	CreateFakeByClient(client6.ID)
	CreateFakeByClient(client7.ID)
	CreateFakeByClient(client8.ID)
	CreateFakeByClient(client9.ID)
	CreateFakeByClient(client10.ID)
	CreateFakeByClient(client11.ID)

	CreateFakeByClinic(clinic2.ID)
	CreateFakeByClinic(clinic1.ID)
	CreateFakeByClinic(clinic3.ID)
}

func CreateFakeByClinic(clinicId uint) {
	//PROMOS
	for i := 1; i < 100; i++ {
		promo := promos.Promo{
			ClinicID: clinicId,
			Title:    "Tu primera eco gratis",
			Text:     "Ven a vernos y consigue que te hagamos la primera eco gratis.",
			Week:     34,
			ImageUrl: "https://s3.eu-central-1.wasabisys.com/stela/weeks/21SM.jpg",
			StartAt:  time.Now().Format("2006-01-02 15:04:05"),
			EndAt:    time.Now().Add(time.Hour * 24 * 7).Format("2006-01-02 15:04:05"),
		}

		database.GormDB.Create(&promo)
	}

}
func CreateFakeByClient(clientId uint) {
	/*
		calculator1 := models.Calculator{ClientID: clientId, Week: 4}
		database.GormDB.Create(&calculator1)

		calculator2 := models.Calculator{ClientID: clientId, Week: 8}
		database.GormDB.Create(&calculator2)
	*/

	// IMAGES ------------------------------------------------------------------------------------------------------------------------------------------------------------------
	for i := 1; i < 20; i++ {
		images := models.Image{
			ClientID: clientId,
			Url:      "https://s3.eu-central-1.wasabisys.com/ecobaby/1/image/MARIANA VICENTE (10).png",
			Size:     0,
		}

		database.GormDB.Create(&images)
	}

	for i := 1; i < 20; i++ {
		videos := models.Video{
			ClientID:     clientId,
			Url:          "https://s3.eu-central-1.wasabisys.com/ecobaby/1/video/MARIANA VICENTE (2) (online-video-cutter.com) copy 10.mp4",
			ThumbnailUrl: "https://s3.eu-central-1.wasabisys.com/ecobaby/1/video/MARIANA VICENTE (2) (online-video-cutter.com) copy 10.mp4-thumb.jpg",
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
			ClientID: clientId,
			Url:      "https://www.youtube.com/watch?v=5qap5aO4i9A",
			Code:     "12345",
		}

		database.GormDB.Create(&streaming)
	}
}
