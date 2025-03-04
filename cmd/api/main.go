package main

import (
	"fmt"

	"github.com/LeoLion02/album-api/config"
	api "github.com/LeoLion02/album-api/internal/api"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
	_ "github.com/google/wire"

	docs "github.com/LeoLion02/album-api/api"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	_, err := config.GetConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	apiV1 := router.Group("api/v1")
	api.RegisterAlbumRoutes(apiV1)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
