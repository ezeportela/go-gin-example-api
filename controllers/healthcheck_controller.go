package controllers

import (
	"time"

	"github.com/ezeportela/meli-challenge/models"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
)

func RegisterHealthCheckController(r *gin.Engine) {
	r.GET("/healthcheck", func(c *gin.Context) {
		dt := time.Now()

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

		err := mgm.Coll(citizen).Create(citizen)

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
