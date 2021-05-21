package models

import "gorm.io/gorm"

type Citizen struct {
	gorm.Model
	Name        string
	Species     string
	Description string
	Weight      int
	Height      int
	PhotoUrl    string
	HasHuman    bool
	Roles       []string
}
