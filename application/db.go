package application

import (
	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/repositories"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDatabase(conf config.Config) {
	err := mgm.SetDefaultConfig(
		nil,
		conf.DatabaseName,
		options.Client().ApplyURI(conf.MongoURI),
	)

	if err != nil {
		panic(err)
	}

	repository := repositories.RoleRepository{}

	repository.Fixtures(conf)
}
