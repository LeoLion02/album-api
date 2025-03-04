package controllers

import (
	"net/http"
	"strconv"

	_ "github.com/LeoLion02/album-api/internal/api/swagger"
	"github.com/LeoLion02/album-api/internal/application/models"
	"github.com/LeoLion02/album-api/internal/application/services"

	"github.com/gin-gonic/gin"
)

type AlbumController struct {
	albumService services.IAlbumService
}

func NewAlbumController(service services.IAlbumService) *AlbumController {
	return &AlbumController{albumService: service}
}

func (controller *AlbumController) GetAll(c *gin.Context) {
	baseLog := InitLog(nil, c.Request.Method+c.Request.RequestURI)
	defer FinishLog(baseLog)

	result := controller.albumService.GetAll(baseLog)

	ReturnResponseFromResult(result, c, baseLog)
}

func (controller *AlbumController) GetById(c *gin.Context) {
	id := c.Param("id")

	baseLog := InitLog(id, c.Request.Method+c.Request.RequestURI)
	defer FinishLog(baseLog)

	albumID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		HandleRequestError(c, "Invalid album ID.", err, baseLog)
		return
	}

	result := controller.albumService.GetById(albumID, baseLog)
	ReturnResponseFromResult(result, c, baseLog)
}

func (controller *AlbumController) Add(c *gin.Context) {
	baseLog := InitLog(nil, c.Request.Method+c.Request.RequestURI)
	defer FinishLog(baseLog)

	var request models.Album
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	baseLog.Request = request

	addResult := controller.albumService.Add(request, baseLog)
	ReturnResponseFromResult(addResult, c, baseLog)
}

func (controller *AlbumController) Update(c *gin.Context) {
	baseLog := InitLog(nil, c.Request.Method+c.Request.RequestURI)
	defer FinishLog(baseLog)

	id := c.Param("id")

	albumId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		HandleRequestError(c, "Invalid album ID.", err, baseLog)
		return
	}

	var request models.Album
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	baseLog.Request = request

	result := controller.albumService.Update(albumId, request, baseLog)
	ReturnResponseFromResult(result, c, baseLog)
}

func (controller *AlbumController) Delete(c *gin.Context) {
	id := c.Param("id")

	baseLog := InitLog(id, c.Request.Method+c.Request.RequestURI)
	defer FinishLog(baseLog)

	albumId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		HandleRequestError(c, "Invalid album ID.", err, baseLog)
		return
	}

	err = controller.albumService.Delete(albumId, baseLog)
	if err != nil {
		HandleInternalError(c, err)
		return
	}

	c.Status(http.StatusOK)
}
