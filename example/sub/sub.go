package sub

import (
	"github.com/mad01/logger"
)

// Run sub
func Run() {
	logger.Info("info")
	logger.Debug("debug")

	context := logger.WithFields(logger.Fields{"context": "Run"})
	context.Error("error")
	context.Debug("debug")
}
