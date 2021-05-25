package controllers

import (
	"net/http"
	"time"

	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/gin-gonic/gin"
)

func HealthcheckHandler(conf config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		dt := time.Now()

		c.JSON(http.StatusOK, gin.H{
			"version":   conf.Version,
			"status":    "OK",
			"timestamp": shared.FormatDateTime(dt),
		})
	}
}
