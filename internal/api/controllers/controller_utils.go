package controllers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/LeoLion02/album-api/internal/application/models"
	"github.com/LeoLion02/album-api/internal/shared/log"

	"github.com/gin-gonic/gin"
)

func ReturnResponseFromResult[T any](result *models.Result[T], c *gin.Context, baseLog *log.BaseLog) {
	if result.Error != nil {
		baseLog.Error = result.Error.Error()
		baseLog.StatusCode = result.HttpStatusCode

		if result.IsInternalError {
			baseLog.Level = slog.LevelError
			baseLog.Response = "An internal error has occurred."
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error has occurred."})
			return
		}

		errorMsg := fmt.Sprintf("An error has occurred: %v", result.Error)
		baseLog.Level = slog.LevelWarn
		baseLog.Response = errorMsg

		c.JSON(result.HttpStatusCode, gin.H{"error": errorMsg})
		return
	}

	baseLog.StatusCode = http.StatusOK
	baseLog.Response = result.Value

	c.IndentedJSON(http.StatusOK, result.Value)
}

func HandleRequestError(c *gin.Context, message string, err error, baseLog *log.BaseLog) {
	baseLog.Error = err.Error()
	baseLog.Response = message
	baseLog.StatusCode = http.StatusBadRequest
	baseLog.Level = slog.LevelWarn

	c.JSON(http.StatusBadRequest, gin.H{"error": message})
}

func HandleInternalError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal error has occured."})
}

func InitLog(request any, endpoint string) *log.BaseLog {
	baseLog := log.NewBaseLog(request, "")
	baseLog.Endpoint = endpoint
	return baseLog
}

func FinishLog(baseLog *log.BaseLog) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	baseLog.Finish()
	logger = logger.With("message", baseLog)
	logger.Log(context.TODO(), baseLog.Level, "OK")
}
