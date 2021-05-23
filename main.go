package main

import (
	"github.com/ezeportela/meli-challenge/application"
)

func main() {
	router := application.SetupRouter()

	router.Run()
}
