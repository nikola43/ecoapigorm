package services

import (
	"errors"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	streamings "github.com/nikola43/ecoapigorm/models/streamings"
	"github.com/nikola43/ecoapigorm/utils"
	"math/rand"
	"strings"
	"time"
)

func GetStreamingByCodeService(code string) (*streamings.Streaming, error) {
	streaming := new(streamings.Streaming)

	err := database.GormDB.Where("code = ?", code).First(&streaming).Error
	if err != nil {
		return streaming, err
	}

	return streaming, nil
}

func CreateStreaming(createStreamingRequest *streamings.CreateStreamingRequest) (*streamings.Streaming, error) {
	streaming := &streamings.Streaming{}
	client := new(models.Client)
	code := ""

	err := database.GormDB.First(&client, createStreamingRequest.ClientID).Error
	if err != nil {
		return nil, err
	}

	err = database.GormDB.Where("url = ?", createStreamingRequest.Url).Find(&streaming).Error
	if err != nil {
		return nil, err
	}
	if streaming.ID > 0 {
		return nil, errors.New("streaming already exist")
	}

	for ok := true; ok; ok = streaming.ID > 0 {
		code = GenerateRandomCode(4)
		err = database.GormDB.Where("code = ?", code).Find(&streaming).Error
		if err != nil {
			return nil, err
		}
	}

	streaming.Url = createStreamingRequest.Url
	streaming.ClientID = createStreamingRequest.ClientID
	streaming.ClinicID = createStreamingRequest.ClinicID
	streaming.Code = code

	err = database.GormDB.Create(&streaming).Error
	if err != nil {
		return streaming, err
	}
	
	sendEmailManager := utils.SendEmailManager{
		ToEmail:               client.Email,
		ToName:                client.Name,
		Template:              "streaming_email.html",
		Subject:               "Nuevo streaming disponible en mimatrona",
	}
	sendEmailManager.SendMail()

	return streaming, nil
}

func GenerateRandomCode(length int) string {
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)

	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}

	return strings.ToUpper(string(code))
}

func DeleteStreamingByID(streamingID uint) error {
	streaming := new(streamings.Streaming)

	err := database.GormDB.First(&streaming, streamingID).Error
	if err != nil {
		return err
	}

	err = database.GormDB.Delete(streaming).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateStreaming(updateStreaming *streamings.Streaming) (*streamings.Streaming, error) {
	streaming := new(streamings.Streaming)

	err := database.GormDB.First(&streaming, updateStreaming.ID).Error
	if err != nil {
		return nil, err
	}

	err = database.GormDB.Model(&streaming).Update("url", updateStreaming.Url).Error
	if err != nil {
		return nil, err
	}

	return streaming, nil
}
