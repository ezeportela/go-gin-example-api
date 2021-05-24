package application

import (
	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/controllers"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupRouter(conf config.Config) *gin.Engine {
	err := mgm.SetDefaultConfig(
		nil,
		conf.DatabaseName,
		options.Client().ApplyURI(conf.MongoURI),
	)

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	controllers.RegisterHealthCheckController(r, conf)
	controllers.RegisterCitizenController(r, conf)

	return r
}
