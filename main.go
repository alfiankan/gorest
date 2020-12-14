package main

import (
	"github.com/alfiankan/gorest/app"
	"github.com/alfiankan/gorest/logger"
)

func main() {
	logger.Info("Starting Application")
	app.Start()
}
