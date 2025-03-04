package data

import (
	"database/sql"

	"github.com/LeoLion02/album-api/config"
	"github.com/LeoLion02/album-api/internal/shared/log"
)

type IDbContext interface {
	Connect(baseLog *log.BaseLog) (*sql.DB, error)
}

type DbContext struct {
	connectionStringConfig config.ConnectionString
}

func NewDbContext() IDbContext {
	appConfig, _ := config.GetConfig()
	return &DbContext{connectionStringConfig: appConfig.ConnectionString}
}

func (dbContext *DbContext) Connect(baseLog *log.BaseLog) (*sql.DB, error) {
	logStep := log.NewLogStep(nil)
	baseLog.AddStep(log.SqlOpenConnection, log.NewLogStep(logStep))

	result, err := sql.Open("sqlserver", dbContext.connectionStringConfig.SqlServer)

	logStep.Finish(&err)

	return result, err
}
