package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ezeportela/meli-challenge/application"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := application.SetupRouter()

	dt := time.Now()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	expected := gin.H{
		"status":    "OK",
		"timestamp": shared.FormatDateTime(dt),
	}

	assert.Equal(t, shared.StringifyInterface(expected), w.Body.String())
}
