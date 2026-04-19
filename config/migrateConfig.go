package config

import "image_storage/models"

func MigrateConfig() {
	DB.AutoMigrate(&models.Image{})
}
