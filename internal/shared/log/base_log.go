package log

import (
	"log/slog"
	"time"
)

type BaseLog struct {
	Request         any
	Response        any
	Endpoint        string
	Level           slog.Level
	StatusCode      int
	Error           string `json:"Error,omitempty"`
	Steps           map[string]*LogStep
	TempoExecucaoMs int64
	timer           time.Time
}

func NewBaseLog(request any, err string) *BaseLog {
	baseLog := &BaseLog{Request: request, Error: err, Level: slog.LevelInfo, Steps: map[string]*LogStep{}}
	baseLog.timer = time.Now()
	return baseLog
}

func (baseLog *BaseLog) AddStep(key string, step *LogStep) {
	baseLog.Steps[key] = step
}

func (baseLog *BaseLog) Finish() {
	baseLog.TempoExecucaoMs = time.Since(baseLog.timer).Milliseconds()
}
