package repositories

import (
	"github.com/ezeportela/meli-challenge/models"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/kamva/mgm/v3"
)

type CitizenRepository struct{}

func (r *CitizenRepository) GetCollection() *mgm.Collection {
	return mgm.Coll(&models.Citizen{})
}

func (r *CitizenRepository) FindById(id string, dst *models.Citizen) error {
	return r.GetCollection().FindByID(id, dst)
}

func (r *CitizenRepository) Update(id string, updates map[string]interface{}) (models.Citizen, error) {
	var citizen models.Citizen
	if err := r.FindById(id, &citizen); err != nil {
		return models.Citizen{}, err
	}

	shared.JsonToStruct(updates, citizen)
	err := r.GetCollection().Update(&citizen)
	return citizen, err
}
