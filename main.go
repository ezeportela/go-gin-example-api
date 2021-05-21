package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/healthcheck", func(c *gin.Context) {
		dt := time.Now()

		c.JSON(200, gin.H{
			"status":    "OK",
			"timestamp": dt.String(),
		})
	})

	r.Run()
}
