package controllers

import (
	"time"

	"github.com/ezeportela/meli-challenge/models"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterHealthCheck(r *gin.Engine) {
	r.GET("/healthcheck", func(c *gin.Context) {
		dt := time.Now()

		err := mgm.SetDefaultConfig(nil, "animalia", options.Client().ApplyURI("mongodb://localhost:27017"))

		if err != nil {
			c.JSON(409, gin.H{
				"err": err,
			})
			return
		}

		citizen := &models.Citizen{
			Name:        "Doge",
			Species:     "dog",
			Description: "A citizen",
			Weight:      30,
			Height:      40,
			PhotoUrl:    "https://",
			HasHuman:    true,
			Roles:       []string{"Civil"},
		}

		err = mgm.Coll(citizen).Create(citizen)

		if err != nil {
			c.JSON(409, gin.H{
				"err": err,
			})
			return
		}

		c.JSON(200, gin.H{
			"status":    "OK",
			"timestamp": dt.String(),
		})
	})
}
