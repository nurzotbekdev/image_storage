package routes

import (
	"image_storage/controllers"
	"image_storage/services"

	"github.com/gin-gonic/gin"
)

func ImageRoutes(r *gin.Engine) {
	imageServices := services.NewImageServices()
	imageController := controllers.NewImageController(imageServices)

	r.POST("/image", imageController.Create)
	r.GET("/image", imageController.AllImage)
	r.PUT("/image/:id", imageController.UpdateImage)
	r.DELETE("/image/:id", imageController.RemoveImage)
}
