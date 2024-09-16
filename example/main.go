package main

import (
	"github.com/KyberNetwork/logger"
)

func main() {
	config := logger.Configuration{
		EnableConsole:    true,
		EnableJSONFormat: false,
		ConsoleLevel:     "info",
		EnableFile:       false,
	}
	log, err := logger.NewLogger(config, logger.LoggerBackendPhuslu)
	if err != nil {
		panic(err)
	}
	log.Info("This is an info message using Phuslu logger")
	log.Debug("This is a debug message using Phuslu logger")
	log.Error("This is an error message using Phuslu logger")
}
