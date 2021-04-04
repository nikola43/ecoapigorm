package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	database "github.com/nikola43/ecoapigorm/database"
	"github.com/nikola43/ecoapigorm/models"
	"github.com/nikola43/ecoapigorm/services"
	"github.com/nikola43/ecoapigorm/utils"
	"github.com/nikola43/ecoapigorm/wasabis3manager"
	"os"
	"strconv"
	"strings"
)

func UploadMultimedia(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)
	clinicId, _ := strconv.ParseUint(context.Params("clinic_id"), 10, 64)
	uploadMode, _ := strconv.ParseUint(context.Params("upload_mode"), 10, 64)
	uploadedFile, err := context.FormFile("file")
	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	if err != nil {
		return context.SendStatus(fiber.StatusUnauthorized)
	}

	bucketName := strings.ToLower(strings.ReplaceAll(employeeTokenClaims.CompanyName, " ", ""))

	clinic := models.Clinic{}
	database.GormDB.First(&clinic, clinicId)
	if clinic.ID < 1 {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	err = services.UploadMultimedia(
		context,
		bucketName,
		clinic.Name,
		uint(clientID),
		uploadedFile,
		uint(uploadMode),
		clinic.ID)
	if err != nil {
		return err
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"file": uploadedFile.Filename,
	})
}

func DownloadAllMultimediaContentByClientID(context *fiber.Ctx) error {
	clientID, _ := strconv.ParseUint(context.Params("client_id"), 10, 64)

	// download images
	images, err := services.GetAllImagesByClientID(uint(clientID))
	os.Mkdir("tempFiles/"+context.Params("client_id")+"/images", os.ModePerm)
	for _, image := range images {

		fmt.Println(image.Url)
		filename := strings.Split(image.Url, "/")[len(strings.Split(image.Url, "/"))-1]
		fmt.Println(filename)
		err = wasabis3manager.WasabiS3Client.DownloadObject(filename, "tempFiles/"+context.Params("client_id")+"/images/"+filename)
		if err != nil {
			fmt.Println(err)
		}
	}
	// download videos
	/*
		videos, err := services.GetAllVideosByClientID(uint(clientID))
		os.Mkdir("tempFiles/"+context.Params("client_id")+"/images", os.ModePerm)
		for _, video := range videos {
			filename := strings.Split(video.Url, "/")[len(strings.Split(video.Url, "/"))-1]
			err = DownloadFile("tempFiles/"+context.Params("client_id")+"/"+filename, video.Url)
			if err != nil {
				fmt.Println(err)
			}
		}
	*/

	return context.SendStatus(fiber.StatusOK)
}

func DeleteHeartbeat(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("id"), 10, 64)
	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	if err != nil {
		return context.SendStatus(fiber.StatusUnauthorized)
	}

	bucketName := strings.ToLower(strings.ReplaceAll(employeeTokenClaims.CompanyName, " ", ""))

	err = services.DeleteHeartbeat(bucketName, uint(id))
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(id)

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func DeleteImage(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("id"), 10, 64)
	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	if err != nil {
		return context.SendStatus(fiber.StatusUnauthorized)
	}

	bucketName := strings.ToLower(strings.ReplaceAll(employeeTokenClaims.CompanyName, " ", ""))

	err = services.DeleteImage(bucketName, uint(id))
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(id)

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func DeleteVideo(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("id"), 10, 64)
	employeeTokenClaims, err := utils.GetEmployeeTokenClaims(context)
	if err != nil {
		return context.SendStatus(fiber.StatusUnauthorized)
	}

	bucketName := strings.ToLower(strings.ReplaceAll(employeeTokenClaims.CompanyName, " ", ""))

	err = services.DeleteVideo(bucketName, uint(id))
	if err != nil {
		return context.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(id)

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}

func DeleteHolographic(context *fiber.Ctx) error {
	id, _ := strconv.ParseUint(context.Params("holopraphic_id"), 10, 64)

	fmt.Println(id)

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
	})
}
