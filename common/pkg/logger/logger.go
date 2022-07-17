package logger

import "go.uber.org/zap"

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
