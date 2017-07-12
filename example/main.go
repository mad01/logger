package main

import (
	"github.com/mad01/logger"
	"github.com/mad01/logger/example/sub"
)

func main() {
	logger.Init(true)
	logger.Info("info msg")
	logger.Debug("debug msg")
	sub.Run()
}
