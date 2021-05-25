package application

import (
	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/controllers"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/gin-gonic/gin"
)

func RegisterHealthCheckController(r *gin.Engine, conf config.Config) {
	r.GET(
		shared.MakeRoute(conf.BasePath, "/healthcheck"),
		controllers.HealthcheckHandler(conf),
	)
}
