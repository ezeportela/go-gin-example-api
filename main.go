package main

import (
	"fmt"

	"github.com/ezeportela/meli-challenge/controllers"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := mgm.SetDefaultConfig(nil, "animalia", options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println("error mongodnb", err)
	}

	r := gin.Default()

	controllers.RegisterHealthCheckController(r)
	controllers.RegisterCitizenController(r)

	r.Run()
}
