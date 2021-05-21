package routes

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterHealthCheck(r *gin.Engine) {
	r.GET("/healthcheck", func(c *gin.Context) {
		dt := time.Now()

		c.JSON(200, gin.H{
			"status":    "OK",
			"timestamp": dt.String(),
		})
	})
}
