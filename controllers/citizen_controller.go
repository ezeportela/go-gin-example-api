package controllers

import (
	"net/http"

	"github.com/ezeportela/meli-challenge/models"
	"github.com/ezeportela/meli-challenge/repositories"
	"github.com/gin-gonic/gin"
)

type CitizenController struct {
	Repository repositories.CitizenRepository
}

// CreateHandler - POST /citizen
func (ctrl *CitizenController) CreateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var citizen models.Citizen

		if err := c.ShouldBindJSON(&citizen); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"data":    nil,
				"message": "Bad Request",
			})
			return
		}

		citizen, err := ctrl.Repository.Insert(citizen)

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
	}
}

// GetByIdHandler - GET /citizen/:id
func (ctrl *CitizenController) GetByIdHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		err := ctrl.Repository.FindById(id, &citizen)

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
	}
}

// UpdateHandler - POST /citizen/:id
func (ctrl *CitizenController) UpdateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		citizen, err := ctrl.Repository.Update(id, updates)
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
	}
}

// DeleteHandler - DELETE /citizen/:id
func (ctrl *CitizenController) DeleteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Params.Get("id")

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "The id is required",
				"data":    nil,
				"message": "Bad Request",
			})
			return
		}

		citizen, err := ctrl.Repository.Delete(id)

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
	}
}

// FilterHandler - POST /citizen/filter
func (ctrl *CitizenController) FilterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		citizens, err := ctrl.Repository.Filter(
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
	}
}

// InsertManyHandler - POST /citizen/batch
func (ctrl *CitizenController) InsertManyHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
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

		result, err := ctrl.Repository.InsertMany(rows)

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
	}
}
