package services

import (
	"errors"
	"fmt"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models/streaming"
	streamings "github.com/nikola43/ecoapigorm/models/streamings"
	"github.com/nikola43/ecoapigorm/utils"
	"math/rand"
	"strings"
	"time"
)

func GetStreamingByCodeService(code string) (streaming.Streaming, error) {
	var streaming = streaming.Streaming{}

	if err := database.GormDB.Where("code = ?", code).
		First(&streaming).Error;

	err != nil {
		return streaming, err
	}

	return streaming, nil
}

func CreateStreaming(createStreamingRequest *streamings.CreateStreamingRequest) (*streaming.Streaming, error) {
	streaming := &streaming.Streaming{}
	code := ""
	fmt.Println(createStreamingRequest)
	for ok := true; ok; ok = streaming.ID > 0 {
		code = GenerateRandomCode(4)
		database.GormDB.Where("code = ?", code).Find(&streaming)
	}
	fmt.Println(code)
	streaming.Url = createStreamingRequest.Url
	streaming.ClientID = createStreamingRequest.ClientID
	streaming.Code = code

	database.GormDB.Create(&streaming)


	return streaming, nil
}

func GenerateRandomCode(length int) string {
	var seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return strings.ToUpper(string(code))
}

func DeleteStreamingByID(streamingID uint) error {
	deleteStreaming := new(streaming.Streaming)

	// todo check clinic is who make action
	// check if employee exist
	utils.GetModelByField(deleteStreaming, "id", streamingID)
	if deleteStreaming.ID < 1 {
		return errors.New("streaming not found")
	}

	// delete employee
	result := database.GormDB.Delete(deleteStreaming)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
