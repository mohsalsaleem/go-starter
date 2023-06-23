package main

import (
	"github.com/mohsalsaleem/go-starter/config"
	"github.com/mohsalsaleem/go-starter/internal/service"
	"github.com/mohsalsaleem/go-starter/logger"
)

func main() {

	// Init config
	config.Init()

	// Init logger
	logger.Init(1)

	logger.Infof("Starting service")

	exampleService := service.New()
	exampleService.Run()

	logger.Infof("Exiting service")
}
