package models

import "github.com/kamva/mgm/v3"

type Citizen struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string   `json:"name" bson:"name" binding:"required"`
	Species          string   `json:"species" bson:"species" binding:"required"`
	Description      string   `json:"description" bson:"description" binding:"required"`
	Weight           int      `json:"weight" bson:"weight" binding:"required"`
	Height           int      `json:"height" bson:"height" binding:"required"`
	PhotoUrl         string   `json:"photo_url" bson:"photo_url" binding:"required"`
	HasHuman         bool     `json:"has_human" bson:"has_human" binding:"required"`
	Roles            []string `json:"roles" bson:"roles" binding:"required"`
}

type CitizenFilter struct {
	Name        string   `json:"name"`
	Species     string   `json:"species"`
	Description string   `json:"description"`
	HasHuman    bool     `json:"has_human"`
	Roles       []string `json:"roles"`
}
