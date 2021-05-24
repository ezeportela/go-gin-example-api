package models

import "github.com/kamva/mgm/v3"

type Role struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" yaml:"name"`
	Asignations      int    `json:"asignations" yaml:"asignations"`
	Exclude          bool   `json:"exclude" yaml:"exclude"`
}
