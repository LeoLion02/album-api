// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package shared

import (
	"github.com/LeoLion02/album-api/internal/api/controllers"
	"github.com/LeoLion02/album-api/internal/application/services"
	"github.com/LeoLion02/album-api/internal/infra/data"
	"github.com/LeoLion02/album-api/internal/infra/data/repositories"
)

// Injectors from wire.go:

func InitializeAlbumController() *controllers.AlbumController {
	iDbContext := data.NewDbContext()
	iAlbumRepository := repositories.NewAlbumRepository(iDbContext)
	iAlbumService := services.NewAlbumService(iAlbumRepository)
	albumController := controllers.NewAlbumController(iAlbumService)
	return albumController
}
