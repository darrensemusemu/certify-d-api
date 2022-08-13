package logger

import (
	"go.uber.org/zap"
)

// Global logger
var Log *Logger

func init() {
	Log, _ = New("global")
	zap.ReplaceGlobals(Log.Desugar())
}

// A shared logger
type Logger struct {
	*zap.SugaredLogger
}

// Creates a new logger
func New(service string) (*Logger, error) {
	config := zap.NewProductionConfig()
	l, err := config.Build()
	if err != nil {
		return nil, err
	}
	lSugar := &Logger{l.Sugar()}
	lSugar.With("service", service)
	return lSugar, nil
}

// TODO: HTTPRequest structured logs
type HTTPRequest struct {
	Logger
}
