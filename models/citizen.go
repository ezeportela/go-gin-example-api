package models

type Citizen struct {
	Id          int
	Name        string
	Species     string
	Description string
	Weight      int
	Height      int
	PhotoUrl    string
	HasHuman    bool
	Roles       []string
}
