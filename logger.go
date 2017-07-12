package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Init func to set current log context
func Init(debug bool) {

	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Only log the warning severity or above.
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// Info log proxy
func Info(args ...interface{}) {
	logrus.Info(args)
}

// Infof log proxy
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args)
}

// Warn log proxy
func Warn(args ...interface{}) {
	logrus.Warn(args)
}

// Warnf log proxy
func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args)
}

// Error log proxy
func Error(args ...interface{}) {
	logrus.Error(args)
}

// Errorf log proxy
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args)
}

// Debug log proxy
func Debug(args ...interface{}) {
	logrus.Debug(args)
}

// Debugf log proxy
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args)
}

// WithFields log proxy
func WithFields(fields logrus.Fields) *logrus.Entry {
	return logrus.WithFields(fields)
}
