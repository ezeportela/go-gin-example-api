package application

import (
	"fmt"

	"github.com/ezeportela/meli-challenge/controllers"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupRouter() *gin.Engine {
	err := mgm.SetDefaultConfig(nil, "animalia", options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println("error mongodb", err)
	}

	r := gin.Default()

	controllers.RegisterHealthCheckController(r)
	controllers.RegisterCitizenController(r)

	return r
}
