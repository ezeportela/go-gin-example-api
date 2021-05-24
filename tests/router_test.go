package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/ezeportela/meli-challenge/application"
	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/models"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestRouter(t *testing.T) {
	conf := config.Config{}
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")
	conf.Setup("../config/default.yml")
	conf.DatabaseName = "animalia_test"

	router := application.SetupRouter(conf)

	mgm.Coll(&models.Citizen{}).DeleteMany(
		context.Background(),
		map[string]interface{}{},
	)

	t.Run("test healthcheck endpoint", func(t *testing.T) {
		dt := time.Now()

		w := CallApi(ApiParams{
			Conf:               conf,
			Method:             "GET",
			Path:               "/healthcheck",
			Test:               t,
			Router:             router,
			ExpectedStatusCode: http.StatusOK,
		})

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

		w := CallApi(ApiParams{
			Conf:               conf,
			Method:             "POST",
			Path:               "/citizen",
			Body:               citizen,
			Test:               t,
			Router:             router,
			ExpectedStatusCode: http.StatusCreated,
		})

		var res map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &res)
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

		w := CallApi(ApiParams{
			Conf:               conf,
			Method:             "GET",
			Path:               url,
			Test:               t,
			Router:             router,
			ExpectedStatusCode: http.StatusOK,
		})

		expected := gin.H{
			"data":  citizen,
			"error": "",
		}

		assert.Equal(t, shared.StringifyInterface(expected), w.Body.String())
	})

	t.Run("test insert citizens endpoint", func(t *testing.T) {

		citizens := map[string]interface{}{
			"data": []map[string]interface{}{
				{
					"name":        "Doge",
					"species":     "dog",
					"description": "the first citizen",
					"weight":      30,
					"height":      40,
					"photo_url":   "https://phantom-marca.unidadeditorial.es/252acdd64f48851f815c16049a789f23/resize/1320/f/jpg/assets/multimedia/imagenes/2021/04/19/16188479459744.jpg",
					"has_human":   false,
					"roles": []string{
						"First Minister",
					},
				},
				{
					"name":        "Kitty",
					"species":     "cat",
					"description": "the second citizen",
					"weight":      5,
					"height":      20,
					"photo_url":   "https://img.webmd.com/dtmcms/live/webmd/consumer_assets/site_images/article_thumbnails/other/cat_relaxing_on_patio_other/1800x1200_cat_relaxing_on_patio_other.jpg",
					"has_human":   true,
					"roles": []string{
						"Treasurer",
					},
				},
				{
					"name":        "Leonardo",
					"species":     "turtle",
					"description": "the third citizen",
					"weight":      2,
					"height":      10,
					"photo_url":   "https://upload.wikimedia.org/wikipedia/commons/thumb/f/f4/Florida_Box_Turtle_Digon3_re-edited.jpg/1200px-Florida_Box_Turtle_Digon3_re-edited.jpg",
					"has_human":   false,
					"roles": []string{
						"General",
					},
				},
			},
		}

		w := CallApi(ApiParams{
			Conf:               conf,
			Method:             "POST",
			Path:               "/citizen/batch",
			Body:               citizens,
			Test:               t,
			Router:             router,
			ExpectedStatusCode: http.StatusCreated,
		})

		var res map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &res)
		assert.NoError(t, err)

		data := res["data"]
		var resMap map[string]interface{}
		shared.CastInterface(data, &resMap)
		insertedIDs := resMap["InsertedIDs"]
		arrInsertedIDs := insertedIDs.([]interface{})

		expected := citizens["data"]
		arrExpected := expected.([]map[string]interface{})

		assert.Equal(t, len(arrExpected), len(arrInsertedIDs))
	})

	t.Run("test filter citizens endpoint", func(t *testing.T) {

		filters := map[string]string{
			"name": "Doge",
		}

		w := CallApi(ApiParams{
			Conf:               conf,
			Method:             "POST",
			Path:               "/citizen/filter",
			Body:               filters,
			Test:               t,
			Router:             router,
			ExpectedStatusCode: http.StatusOK,
		})

		var res map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &res)
		assert.NoError(t, err)

		data := res["data"]
		citizens := data.([]interface{})

		assert.GreaterOrEqual(t, len(citizens), 1)
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

		w := CallApi(ApiParams{
			Conf:               conf,
			Method:             "POST",
			Path:               url,
			Body:               updates,
			Test:               t,
			Router:             router,
			ExpectedStatusCode: http.StatusOK,
		})

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

		w := CallApi(ApiParams{
			Conf:               conf,
			Method:             "DELETE",
			Path:               url,
			Test:               t,
			Router:             router,
			ExpectedStatusCode: http.StatusOK,
		})

		expected := gin.H{
			"message": "The citizen has been deleted successfully",
			"data":    citizen,
			"error":   "",
		}

		assert.Equal(t, shared.StringifyInterface(expected), w.Body.String())
	})

}
