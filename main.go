package main

import (
	"github.com/ezeportela/meli-challenge/application"
	"github.com/ezeportela/meli-challenge/config"
)

func main() {
	conf := config.Config{}
	conf.Setup("./config/default.yml")

	application.SetupDatabase(conf)
	router := application.SetupRouter(conf)

	router.Run()
}
