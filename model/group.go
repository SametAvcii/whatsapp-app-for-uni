package models

import "github.com/kamva/mgm/v3"

type Group struct {
	mgm.DefaultModel `bson:",inline" json:"id"`
	Name             string `bson:"group_name" json:"group_name"`
	DepartmentCode   uint16 `json:"department_code" bson:"department_code"`
	Link             string `json:"group_link" bson:"group_link"`
	IsVerified       bool   `json:"is_verified" bson:"is_verified"`
}

func NewGroup(group *Group) *Group {
	return &Group{
		
		Name:           group.Name,
		DepartmentCode: group.DepartmentCode,
		Link:           group.Link,
		IsVerified:     group.IsVerified,
	}
}
