package main

import (
	"banners_oto/config"
	"banners_oto/internal/app"
	"fmt"

	"go.uber.org/zap"
)

func main() {

	logger := zap.Must(zap.NewProduction())
	defer func() {
		if err := logger.Sync(); err != nil {
			logger.Warn(fmt.Sprintf("Logger syncing is'nt available -> %v", err))
		}
	}()

	config.Read(logger)
	app.Run(logger)
	logger.Info("The server has shut down ... ")

}
