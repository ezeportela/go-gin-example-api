package tests

import (
	"os"
	"testing"

	"github.com/ezeportela/meli-challenge/application"
	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/models"
	"github.com/ezeportela/meli-challenge/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCheckRoles(t *testing.T) {
	conf := config.Config{}
	os.Setenv("MONGO_URI", "mongodb://localhost:27017")
	conf.Setup("../config/default.yml")
	conf.DatabaseName = "animalia_test"

	application.SetupDatabase(conf)
	repository := repositories.CitizenRepository{}

	citizen := models.Citizen{
		Name:        "Doge",
		Species:     "dog",
		Description: "A citizen",
		Weight:      30,
		Height:      40,
		PhotoUrl:    "https://",
		HasHuman:    true,
		Roles:       []string{"First Minister"},
	}

	ok, err := repository.CheckRoles(citizen)

	assert.NoError(t, err)
	assert.False(t, ok)
}
