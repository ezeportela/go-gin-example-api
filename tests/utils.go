package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ApiParams struct {
	Conf               config.Config
	Method             string
	Path               string
	Body               interface{}
	Test               *testing.T
	Router             *gin.Engine
	ExpectedStatusCode int
}

func CallApi(params ApiParams) *httptest.ResponseRecorder {
	reqBody := new(bytes.Buffer)

	if params.Body != nil {
		body, err := json.Marshal(params.Body)
		reqBody = bytes.NewBuffer(body)
		assert.NoError(params.Test, err)
	}

	url := shared.MakeRoute(params.Conf.BasePath, params.Path)

	req, err := http.NewRequest(
		params.Method,
		url,
		reqBody,
	)

	assert.NoError(params.Test, err)

	w := httptest.NewRecorder()
	params.Router.ServeHTTP(w, req)
	assert.Equal(params.Test, params.ExpectedStatusCode, w.Code)

	return w
}
