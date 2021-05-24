package repositories

import (
	"context"

	"github.com/ezeportela/meli-challenge/config"
	"github.com/ezeportela/meli-challenge/models"
	"github.com/kamva/mgm/v3"
)

type RoleRepository struct{}

func (r *RoleRepository) GetColl() *mgm.Collection {
	return mgm.Coll(&models.Role{})
}

func (r *RoleRepository) Fixtures(conf config.Config) {
	count, err := r.GetColl().CountDocuments(context.Background(), map[string]interface{}{})

	if err != nil {
		panic(err)
	}

	if count == 0 {
		rows := make([]interface{}, len(conf.Roles))
		for i, item := range conf.Roles {
			rows[i] = item
		}

		r.GetColl().InsertMany(context.Background(), rows)
	}
}

func (r *RoleRepository) FindById(id string, dst models.Role) error {
	return r.GetColl().FindByID(id, &dst)
}
