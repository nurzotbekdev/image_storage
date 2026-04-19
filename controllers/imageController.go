package controllers

import (
	"errors"
	"image_storage/helper"
	"image_storage/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageController struct {
	ImageService services.ImageServices
}

func NewImageController(image services.ImageServices) *ImageController {
	return &ImageController{ImageService: image}
}

func (image *ImageController) Create(ctx *gin.Context) {
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Image is required",
		})
		return
	}

	_, err = image.ImageService.Upload(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to upload image",
			"detail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Image added to database",
	})
}

func (image *ImageController) AllImage(ctx *gin.Context) {
	images, err := image.ImageService.GetAllImage()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal server error",
		})
		return
	}

	response := make([]gin.H, 0)
	for _, row := range images {
		response = append(response, gin.H{
			"id":         row.ID,
			"image":      row.Image,
			"created_at": row.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

func (image *ImageController) UpdateImage(ctx *gin.Context) {
	imageID, err := helper.ParseUintParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Image is required",
		})
		return
	}

	if err := image.ImageService.EditImage(imageID, file); err != nil {
		if errors.Is(err, services.ImageNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": services.ImageNotFound.Error(),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Image update to database",
	})
}

func (image *ImageController) RemoveImage(ctx *gin.Context) {
	imageID, err := helper.ParseUintParam(ctx, "id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := image.ImageService.DeleteImage(imageID); err != nil {
		if errors.Is(err, services.ImageNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": services.ImageNotFound.Error(),
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Image delete to database",
	})
}
