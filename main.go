package main

import (
	"github.com/ViniciusMAlmeida/goopportunities/config"
	"github.com/ViniciusMAlmeida/goopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()
}
