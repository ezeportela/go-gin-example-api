package controllers

import (
	"net/http"

	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/models"
	"github.com/ezeportela/meli-challenge/repositories"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/gin-gonic/gin"
)

func RegisterCitizenController(r *gin.Engine, conf config.Config) {
	repository := repositories.CitizenRepository{}

	// Insert - POST /citizen
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

		citizen, err := repository.Insert(citizen)

		if err != nil {
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

	// Get by Id - GET /citizen/:id
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

		err := repository.FindById(id, &citizen)

		if err != nil {
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

	// Update - POST /citizen/:id
	r.POST(shared.MakeRoute(conf.BasePath, "/citizen/:id"), func(c *gin.Context) {
		var updates map[string]interface{}

		id, ok := c.Params.Get("id")

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "The id is required",
				"data":    nil,
				"message": "Bad Request",
			})
			return
		}

		if err := c.ShouldBindJSON(&updates); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "Bad Request",
			})
			return
		}

		citizen, err := repository.Update(id, updates)
		if err != nil {
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

	// Delete - DELETE /citizen/:id
	r.DELETE(shared.MakeRoute(conf.BasePath, "/citizen/:id"), func(c *gin.Context) {
		id, ok := c.Params.Get("id")

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "The id is required",
				"data":    nil,
				"message": "Bad Request",
			})
			return
		}

		citizen, err := repository.Delete(id)

		if err != nil {
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

	// Filter citizens - /citizen/filter
	r.POST(shared.MakeRoute(conf.BasePath, "/citizen/filter"), func(c *gin.Context) {
		var params map[string]interface{}

		if err := c.ShouldBindJSON(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "Bad Request",
			})
			return
		}

		var limit int64 = 50
		citizens, err := repository.Filter(
			params,
			limit,
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

	// Insert Many - POST /citizen/batch
	r.POST(shared.MakeRoute(conf.BasePath, "/citizen/batch"), func(c *gin.Context) {
		var body map[string]interface{}

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

		result, err := repository.InsertMany(rows)

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
