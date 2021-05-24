package application

import (
	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(conf config.Config) *gin.Engine {
	r := gin.Default()

	controllers.RegisterHealthCheckController(r, conf)
	controllers.RegisterCitizenController(r, conf)

	return r
}
