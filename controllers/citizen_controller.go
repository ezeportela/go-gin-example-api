package controllers

import (
	"context"
	"net/http"

	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/models"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterCitizenController(r *gin.Engine, conf config.Config) {
	r.POST(shared.MakeRoute(conf.BasePath, "/citizen"), func(c *gin.Context) {
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
				"message": "The citizen has not been created"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data":    citizen,
			"error":   "",
			"message": "The citizen has been created successfully",
		})
	})

	r.GET(shared.MakeRoute(conf.BasePath, "/citizen/:id"), func(c *gin.Context) {
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
				"message": "The citizen has not been found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  citizen,
			"error": "",
		})
	})

	r.POST(shared.MakeRoute(conf.BasePath, "/citizen/:id"), func(c *gin.Context) {
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
				"message": "The citizen has not been found"})
			return
		}

		if err := c.ShouldBindJSON(&citizen); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "Bad Request",
			})
			return
		}

		if err := mgm.Coll(&citizen).Update(&citizen); err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "The citizen has not been updated"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    citizen,
			"error":   "",
			"message": "The citizen has been updated successfully",
		})
	})

	r.DELETE(shared.MakeRoute(conf.BasePath, "/citizen/:id"), func(c *gin.Context) {
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
				"message": "The citizen has not been found"})
			return
		}

		if err := mgm.Coll(&citizen).Delete(&citizen); err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "The citizen has not been deleted"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    citizen,
			"error":   "",
			"message": "The citizen has been deleted successfully",
		})
	})

	r.POST(shared.MakeRoute(conf.BasePath, "/citizen/filter"), func(c *gin.Context) {
		var params map[string]interface{}
		var citizen models.Citizen
		var citizens []models.Citizen

		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "Bad Request",
			})
			return
		}

		var limit int64 = 50
		err := mgm.Coll(&citizen).SimpleFind(
			&citizens,
			params,
			&options.FindOptions{Limit: &limit},
		)

		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "The citizens have not been found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":  citizens,
			"error": "",
		})
	})

	r.POST(shared.MakeRoute(conf.BasePath, "/citizen/batch"), func(c *gin.Context) {
		var body map[string]interface{}
		// var citizens []models.Citizen

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "Bad Request",
			})
			return
		}

		data := body["data"]
		rows := data.([]interface{})

		result, err := mgm.CollectionByName("citizens").InsertMany(
			context.Background(),
			rows,
		)

		if err != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "The citizens have not been created"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data":    result,
			"error":   "",
			"message": "The citizens have been created successfully",
		})
	})
}
