package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ezeportela/meli-challenge/application"
	"github.com/ezeportela/meli-challenge/models"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestRouter(t *testing.T) {
	router := application.SetupRouter()

	t.Run("test healthcheck endpoint", func(t *testing.T) {
		dt := time.Now()
		req, err := http.NewRequest("GET", "/healthcheck", nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)

		expected := gin.H{
			"status":    "OK",
			"timestamp": shared.FormatDateTime(dt),
		}

		assert.Equal(t, shared.StringifyInterface(expected), w.Body.String())
	})

	t.Run("test create citizen endpoint", func(t *testing.T) {
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

		body, err := json.Marshal(citizen)
		b := bytes.NewBuffer(body)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", "/citizen", b)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		var res map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &res)
		assert.NoError(t, err)

		data := res["data"]
		var resCitizen models.Citizen
		shared.CastInterface(data, &resCitizen)
		citizen.ID = resCitizen.ID
		citizen.CreatedAt = resCitizen.CreatedAt
		citizen.UpdatedAt = resCitizen.UpdatedAt

		expected := gin.H{
			"message": "The citizen has been created successfully",
			"data":    citizen,
			"error":   "",
		}

		assert.Equal(t, shared.StringifyInterface(expected), w.Body.String())
	})

	t.Run("test get by id citizen endpoint", func(t *testing.T) {

		var citizen models.Citizen
		err := mgm.Coll(&citizen).First(bson.M{}, &citizen)

		assert.NoError(t, err)

		url := fmt.Sprintf("/citizen/%s", citizen.ID.Hex())
		req, err := http.NewRequest("GET", url, nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		expected := gin.H{
			"message": "citizen data",
			"data":    citizen,
			"error":   "",
		}

		assert.Equal(t, shared.StringifyInterface(expected), w.Body.String())
	})

	t.Run("test update citizen endpoint", func(t *testing.T) {

		var citizen models.Citizen
		err := mgm.Coll(&citizen).First(bson.M{}, &citizen)

		assert.NoError(t, err)

		url := fmt.Sprintf("/citizen/%s", citizen.ID.Hex())

		updates := map[string]string{
			"name":    "Kitty",
			"species": "cat",
		}

		body, err := json.Marshal(updates)
		b := bytes.NewBuffer(body)
		assert.NoError(t, err)

		req, err := http.NewRequest("POST", url, b)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		var res map[string]interface{}
		err = json.Unmarshal(w.Body.Bytes(), &res)
		assert.NoError(t, err)

		data := res["data"]
		var resCitizen models.Citizen
		shared.CastInterface(data, &resCitizen)
		citizen.Name = resCitizen.Name
		citizen.Species = resCitizen.Species
		citizen.CreatedAt = resCitizen.CreatedAt
		citizen.UpdatedAt = resCitizen.UpdatedAt

		expected := gin.H{
			"message": "The citizen has been updated successfully",
			"data":    citizen,
			"error":   "",
		}

		assert.Equal(t, shared.StringifyInterface(expected), w.Body.String())
	})

	t.Run("test delete citizen endpoint", func(t *testing.T) {

		var citizen models.Citizen
		err := mgm.Coll(&citizen).First(bson.M{}, &citizen)

		assert.NoError(t, err)

		url := fmt.Sprintf("/citizen/%s", citizen.ID.Hex())
		fmt.Println(url)

		req, err := http.NewRequest("DELETE", url, nil)
		assert.NoError(t, err)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		expected := gin.H{
			"message": "The citizen has been deleted successfully",
			"data":    citizen,
			"error":   "",
		}

		assert.Equal(t, shared.StringifyInterface(expected), w.Body.String())
	})

}
