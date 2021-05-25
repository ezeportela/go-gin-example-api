package repositories

import (
	"context"
	"errors"

	underscore "github.com/ahl5esoft/golang-underscore"
	"github.com/ezeportela/meli-challenge/models"
	"github.com/ezeportela/meli-challenge/shared"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
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

func (r *CitizenRepository) CheckRoles(model models.Citizen) (bool, error) {
	roleRepository := RoleRepository{}
	roles, err := roleRepository.GetAll()

	if err != nil {
		return false, err
	}

	checked := false

	for _, role := range model.Roles {
		var value models.Role
		underscore.Chain(roles).FindBy(map[string]interface{}{
			"name": role,
		}).Value(&value)

		if value.Asignations > 0 {
			count, err := r.GetColl().CountDocuments(context.Background(), map[string]interface{}{
				"roles": value.Name,
				"_id":   bson.M{"$ne": model.ID},
			})

			if err != nil {
				return false, err
			}

			if count > 0 {
				return false, nil
			}
		}

		if value.Exclude && len(model.Roles) > 1 {
			return false, nil
		}

		if value.Name != "" {
			checked = true
		}
	}

	return checked, nil
}

func (r *CitizenRepository) Insert(model models.Citizen) (models.Citizen, error) {
	ok, _ := r.CheckRoles(model)

	if !ok {
		return model, errors.New("the citizen cannot be inserted: the roles are incorrect")
	}

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

	shared.JsonToStruct(updates, &citizen)

	ok, _ := r.CheckRoles(citizen)

	if !ok {
		return citizen, errors.New("the citizen cannot be updated: the roles are incorrect")
	}

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
