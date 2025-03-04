package log

import "time"

type LogStep struct {
	Request         any
	TempoExecucaoMs int64
	Error           string `json:"Error,omitempty"`
	timer           time.Time
}

func NewLogStep(request any) *LogStep {
	logStep := &LogStep{Request: request}
	logStep.timer = time.Now()
	return logStep
}

func (logStep *LogStep) Finish(err *error) {
	logStep.TempoExecucaoMs = time.Since(logStep.timer).Milliseconds()

	if logStep.Error == "" && err != nil && *err != nil {
		logStep.Error = (*err).Error()
	}
}
