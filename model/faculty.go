package models

import "github.com/kamva/mgm/v3"

type Faculty struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `bson:"name" json:"name"`
	Code             uint16 `bson:"code" json:"code"`
}

func NewFaculty(faculty *Faculty) *Faculty {
	return &Faculty{
		Name: faculty.Name,
		Code: faculty.Code,
	}
}
