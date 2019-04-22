package logging

import (
	"go.uber.org/zap"
)

// Logger is where all logs can be sent, fast and friendly
var Logger *zap.SugaredLogger

// Init starts the Logger instance
func Init() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	Logger = logger.Sugar()
	return
}
