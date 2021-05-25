package application

import (
	"github.com/ezeportela/meli-challenge/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter(conf config.Config) *gin.Engine {
	r := gin.Default()

	RegisterHealthCheckController(r, conf)
	RegisterCitizenController(r, conf)

	return r
}
