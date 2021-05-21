package main

import (
	"github.com/ezeportela/meli-challenge/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	controllers.RegisterHealthCheck(r)

	r.Run()
}
