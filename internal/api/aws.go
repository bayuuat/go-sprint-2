package api

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bayuuat/go-sprint-2/dto"
	"github.com/bayuuat/go-sprint-2/internal/middleware"
	"github.com/bayuuat/go-sprint-2/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type awsApi struct{}

func NewAws(app *fiber.App) {

	da := awsApi{}

	user := app.Group("/v1/file")

	user.Use(middleware.JWTProtected)
	user.Post("/", da.UploadFile)
}

func (da awsApi) UploadFile(ctx *fiber.Ctx) error {
	_, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(400).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: err.Error()})
	}
	defer file.Close()

	fileName := fileHeader.Filename
	fileExt := strings.ToLower(filepath.Ext(fileName))

	allowedExtensions := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}

	if !allowedExtensions[fileExt] {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: "Only jpeg, jpg, and png files are allowed"})
	}

	if fileHeader.Size > 100*1024 {
		return ctx.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Message: "File size exceeds 100KiB"})
	}

	bucketName := os.Getenv("AWS_BUCKET_NAME")

	bucket := bucketName
	prefix := "images"

	res, err := utils.UploadFileToS3(file, fileHeader, bucket, prefix)
	if err != nil {
		log.Println("Error uploading file:", err)
		return ctx.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"uri": res})
}
