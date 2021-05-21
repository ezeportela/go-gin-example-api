package main

import (
	"github.com/ezeportela/meli-challenge/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.RegisterHealthCheck(r)

	r.Run()
}
