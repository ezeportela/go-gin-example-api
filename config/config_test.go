package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")

	config := Config{}

	err := config.Setup("../config/default.yml")

	assert.NoError(t, err)

	assert.NotEmpty(t, config.MongoURI)
}
