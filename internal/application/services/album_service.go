package services

import (
	"errors"
	"net/http"

	"github.com/LeoLion02/album-api/internal/application/models"
	entities "github.com/LeoLion02/album-api/internal/domain/entities"
	"github.com/LeoLion02/album-api/internal/infra/data/repositories"
	logModels "github.com/LeoLion02/album-api/internal/shared/log"
)

type IAlbumService interface {
	GetAll(baseLog *logModels.BaseLog) *models.Result[[]models.Album]
	GetById(id int64, baseLog *logModels.BaseLog) *models.Result[*models.Album]
	Add(albumRequest models.Album, baseLog *logModels.BaseLog) *models.Result[*models.Album]
	Update(id int64, request models.Album, baseLog *logModels.BaseLog) *models.Result[*models.Album]
	Delete(id int64, baseLog *logModels.BaseLog) error
}

type AlbumService struct {
	albumRepository repositories.IAlbumRepository
}

func NewAlbumService(albumRepository repositories.IAlbumRepository) IAlbumService {
	return &AlbumService{albumRepository: albumRepository}
}

func (albumService *AlbumService) GetAll(baseLog *logModels.BaseLog) *models.Result[[]models.Album] {
	albums, err := albumService.albumRepository.FindAll(baseLog)
	if err != nil {
		return models.NewResultFailure[[]models.Album](err, true, http.StatusInternalServerError)
	}

	if albums == nil || len(*albums) == 0 {
		return models.NewResultSuccess([]models.Album{})
	}

	albumsResponse := make([]models.Album, 0, len(*albums))
	for _, album := range *albums {
		albumResponse := models.NewAlbum(album.ID, album.Title, album.Artist, album.Price)
		albumsResponse = append(albumsResponse, *albumResponse)
	}

	return models.NewResultSuccess(albumsResponse)
}

func (albumService *AlbumService) GetById(id int64, baseLog *logModels.BaseLog) *models.Result[*models.Album] {
	album, err := albumService.albumRepository.FindById(id, baseLog)

	if err != nil {
		return models.NewResultFailure[*models.Album](err, true, http.StatusInternalServerError)
	}

	if album == nil {
		return models.NewResultFailure[*models.Album](errors.New("album not found"), false, http.StatusNotFound)
	}

	response := models.NewAlbum(album.ID, album.Title, album.Artist, album.Price)
	return models.NewResultSuccess(response)
}

func (albumService *AlbumService) Add(albumRequest models.Album, baseLog *logModels.BaseLog) *models.Result[*models.Album] {
	album := &entities.Album{
		Title:  albumRequest.Title,
		Artist: albumRequest.Artist,
		Price:  albumRequest.Price,
	}

	newAlbumId, err := albumService.albumRepository.Add(album, baseLog)
	if err != nil {
		return models.NewResultFailure[*models.Album](err, true, http.StatusInternalServerError)
	}

	albumRequest.ID = *newAlbumId
	return models.NewResultSuccess(&albumRequest)
}

func (albumService *AlbumService) Update(id int64, request models.Album, baseLog *logModels.BaseLog) *models.Result[*models.Album] {
	album, err := albumService.albumRepository.FindById(id, baseLog)
	if err != nil {
		return models.NewResultFailure[*models.Album](err, true, http.StatusInternalServerError)
	}

	if album == nil {
		return models.NewResultFailure[*models.Album](errors.New("album not found"), false, http.StatusBadRequest)
	}

	album.Title = request.Title
	album.Artist = request.Artist
	album.Price = request.Price

	err = albumService.albumRepository.Update(album, baseLog)
	if err != nil {
		return models.NewResultFailure[*models.Album](err, true, http.StatusInternalServerError)
	}

	return models.NewResultSuccess[*models.Album](nil)
}

func (albumService *AlbumService) Delete(id int64, baseLog *logModels.BaseLog) error {
	return albumService.albumRepository.Delete(id, baseLog)
}
