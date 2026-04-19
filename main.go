package main

import (
	"image_storage/config"
	"image_storage/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.EnvConfig()
	config.DatabaseConfig()
	config.MigrateConfig()
	config.InitSupabase()

	router := gin.Default()

	routes.ImageRoutes(router)

	router.Run()
}
