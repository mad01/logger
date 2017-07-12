package sub

import (
	"github.com/mad01/logger"
	"github.com/sirupsen/logrus"
)

// Run sub
func Run() {
	logger.Info("sub: info")
	logger.Debug("sub: debug")

	context := logger.WithFields(logrus.Fields{"func": "sub.Run"})
	context.Error("sub: error")
	context.Debug("sub: debug")
}
