package main

import (
	"github.com/edubarr/gojobs/config"
	"github.com/edubarr/gojobs/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Init configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// Init router
	router.Initialize()
}
