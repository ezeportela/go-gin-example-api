package repositories

import (
	"context"

	"github.com/ezeportela/meli-challenge/models"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CitizenRepository struct {
}

func (r *CitizenRepository) GetColl() *mgm.Collection {
	return mgm.Coll(&models.Citizen{})
}

func (r *CitizenRepository) FindById(id string, dst *models.Citizen) error {
	return r.GetColl().FindByID(id, dst)
}

func (r *CitizenRepository) Insert(model models.Citizen) (models.Citizen, error) {
	err := r.GetColl().Create(&model)
	return model, err
}

func (r *CitizenRepository) InsertMany(data []interface{}) (*mongo.InsertManyResult, error) {
	return r.GetColl().InsertMany(
		context.Background(),
		data,
	)
}

func (r *CitizenRepository) Update(id string, updates map[string]interface{}) (models.Citizen, error) {
	var citizen models.Citizen
	if err := r.FindById(id, &citizen); err != nil {
		return models.Citizen{}, err
	}

	shared.JsonToStruct(updates, citizen)
	err := r.GetColl().Update(&citizen)
	return citizen, err
}

func (r *CitizenRepository) Delete(id string) (models.Citizen, error) {
	var citizen models.Citizen
	if err := r.FindById(id, &citizen); err != nil {
		return models.Citizen{}, err
	}

	err := r.GetColl().Delete(&citizen)
	return citizen, err
}

func (r *CitizenRepository) Filter(params interface{}, limit int64) ([]models.Citizen, error) {
	var citizens []models.Citizen
	err := r.GetColl().SimpleFind(
		&citizens,
		params,
		&options.FindOptions{Limit: &limit},
	)

	if err != nil {
		return []models.Citizen{}, err
	}

	return citizens, nil
}
