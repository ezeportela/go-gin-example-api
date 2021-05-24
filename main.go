package main

import (
	"github.com/ezeportela/meli-challenge/application"
	"github.com/ezeportela/meli-challenge/config"
)

func main() {
	config := config.Config{}
	config.Setup("./config/default.yml")

	router := application.SetupRouter()

	router.Run()
}
