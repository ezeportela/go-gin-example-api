package controllers

import (
	"net/http"
	"time"

	"github.com/ezeportela/meli-challenge/shared"
	"github.com/gin-gonic/gin"
)

func RegisterHealthCheckController(r *gin.Engine) {
	r.GET("/healthcheck", func(c *gin.Context) {
		dt := time.Now()

		c.JSON(http.StatusOK, gin.H{
			"status":    "OK",
			"timestamp": shared.FormatDateTime(dt),
		})
	})
}
