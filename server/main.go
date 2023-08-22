package main

import (
	"github.com/tiburciohugo/go-react-todo/config"
	"github.com/tiburciohugo/go-react-todo/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Initialize()
	if err != nil {
		logger.Errorf("error initializing config: %v", err.Error())
		return
	}
	router.Initialize()
}
