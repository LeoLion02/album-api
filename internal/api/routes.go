package api

import (
	"github.com/LeoLion02/album-api/internal/shared"

	"github.com/gin-gonic/gin"
)

func RegisterAlbumRoutes(group *gin.RouterGroup) {
	albumController := shared.InitializeAlbumController()
	group.GET("/album", albumController.GetAll)
	group.GET("/album/:id", albumController.GetById)
	group.POST("/album", albumController.Add)
	group.PUT("/album/:id", albumController.Update)
	group.DELETE("/album/:id", albumController.Delete)
}
