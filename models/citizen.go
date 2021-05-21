package models

import "github.com/kamva/mgm/v3"

type Citizen struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string   `json:"name" bson:"name"`
	Species          string   `json:"species" bson:"species"`
	Description      string   `json:"description" bson:"description"`
	Weight           int      `json:"weight" bson:"weight"`
	Height           int      `json:"height" bson:"height"`
	PhotoUrl         string   `json:"photo_url" bson:"photo_url"`
	HasHuman         bool     `json:"has_human" bson:"has_human"`
	Roles            []string `json:"roles" bson:"roles"`
}
