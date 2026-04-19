package services

import (
	"errors"
	"image_storage/config"
	"image_storage/helper"
	"image_storage/models"
	"mime/multipart"
)

type ImageServices interface {
	Upload(fileHeader *multipart.FileHeader) (string, error)
	GetAllImage() ([]models.Image, error)
	EditImage(ID uint, fileHeader *multipart.FileHeader) error
	DeleteImage(ID uint) error
}

type imageServices struct{}

func NewImageServices() ImageServices {
	return &imageServices{}
}

var (
	ImageNotFound = errors.New("Image not found")
)

func (s *imageServices) Upload(fileHeader *multipart.FileHeader) (string, error) {
	url, err := helper.UploadImage(fileHeader)
	if err != nil {
		return "", err
	}

	image := models.Image{Image: url}
	if err := config.DB.Create(&image).Error; err != nil {
		return "", err
	}

	return url, nil
}

func (s *imageServices) GetAllImage() ([]models.Image, error) {
	var images []models.Image
	err := config.DB.Find(&images).Error
	return images, err
}

func (s *imageServices) EditImage(ID uint, fileHeader *multipart.FileHeader) error {
	var image models.Image
	if err := config.DB.First(&image, ID).Error; err != nil {
		return ImageNotFound
	}

	if err := helper.DeleteImage(image.Image); err != nil {
		return err
	}

	url, err := helper.UploadImage(fileHeader)
	if err != nil {
		return err
	}

	image.Image = url
	if err := config.DB.Save(&image).Error; err != nil {
		return err
	}

	return nil
}

func (s *imageServices) DeleteImage(ID uint) error {
	var image models.Image
	if err := config.DB.First(&image, ID).Error; err != nil {
		return ImageNotFound
	}
	if err := helper.DeleteImage(image.Image); err != nil {
		return err
	}
	if err := config.DB.Unscoped().Delete(&image).Error; err != nil {
		return err
	}

	return nil
}
