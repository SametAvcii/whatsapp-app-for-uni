package models

import "github.com/kamva/mgm/v3"

type Department struct {
	mgm.DefaultModel `bson:",inline"`
	FacultyID        uint16 `bson:"faculty_id" json:"faculty_id"`
	Name             string `bson:"department_name" json:"department_name"`
	Code             uint16 `bson:"code" json:"code"`
}

func NewDepartment(department *Department) *Department {
	return &Department{
		FacultyID: department.FacultyID,
		Name:      department.Name,
		Code:      department.Code,
	}
}
