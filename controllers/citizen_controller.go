package controllers

import (
	"net/http"

	"github.com/ezeportela/meli-challenge/models"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
)

func RegisterCitizenController(r *gin.Engine) {
	r.POST("/citizen", func(c *gin.Context) {
		var citizen models.Citizen

		if err := c.ShouldBindJSON(&citizen); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "Bad Request",
			})
			return
		}

		if err := mgm.Coll(&citizen).Create(&citizen); err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "The Citizen has not been created successfully"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data":    citizen,
			"error":   "",
			"message": "The Citizen has been created successfully",
		})
	})

	r.GET("/citizen/:id", func(c *gin.Context) {
		var citizen models.Citizen

		id, ok := c.Params.Get("id")

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "The id is required",
				"data":    nil,
				"message": "Bad Request",
			})
			return
		}

		if err := mgm.Coll(&citizen).FindByID(id, &citizen); err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "The Citizen has not been found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    citizen,
			"error":   "",
			"message": "Citizen data",
		})
	})
}
