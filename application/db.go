package application

import (
	"context"

	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/models"
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

	count, err := mgm.Coll(&models.Role{}).CountDocuments(context.Background(), map[string]interface{}{})

	if err != nil {
		panic(err)
	}

	if count == 0 {
		rows := make([]interface{}, len(conf.Roles))
		for i, item := range conf.Roles {
			rows[i] = item
		}

		mgm.Coll(&models.Role{}).InsertMany(context.Background(), rows)
	}
}
