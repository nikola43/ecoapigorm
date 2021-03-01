package controllers

import (
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/utils"
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func Migrate()  {
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
}

func CreateFakeData() {
	user := models.Client{Name: "Paulo", Email: "pauloxti@gmail.com", Password: utils.HashPassword([]byte("paulo"))}
	GormDB.Create(&user)

	userMigue := models.Client{Name: "Migue", Email: "migue@gmail.com", Password: utils.HashPassword([]byte("migue"))}
	GormDB.Create(&userMigue)

	images := make([]models.Image, 0)
	images = append(images, models.Image{
		ClientID: 2,
		Url:      "https://www.miecografia4d.com/wp-content/uploads/Ecografia-4-d-oferta-44.jpg",
		Size:     0,
	})

	images = append(images, models.Image{
		ClientID: 2,
		Url:      "https://www.miecografia4d.com/wp-content/uploads/Ecografia-4-d-oferta-44.jpg",
		Size:     0,
	})

	GormDB.Create(&images)

	videos := make([]models.Video, 0)
	videos = append(videos, models.Video{
		ClientID:     2,
		Url:          "https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
		ThumbnailUrl: "https://www.miecografia4d.com/wp-content/uploads/Ecografia-4-d-oferta-44.jpg",
		Size:         0,
	})

	videos = append(videos, models.Video{
		ClientID:     2,
		Url:          "https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
		ThumbnailUrl: "https://www.miecografia4d.com/wp-content/uploads/Ecografia-4-d-oferta-44.jpg",
		Size:         0,
	})

	GormDB.Create(&videos)
}
