package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Password         string `json:"password" bson:"password"`
	SchoolID         int32  `json:"school_id" bson:"school_id"`
	Email            string `json:"email" bson:"email"`
	Verified         bool   `json:"verified" bson:"verified"`
	ApiKey           string `json:"api_key" bson:"api_key"`
	IsAdmin          bool   `json:"is_admin" bson:"is_admin"`
}

func NewUser(user *User) *User {
	return &User{
		Name:     user.Name,
		SchoolID: user.SchoolID,
		Email:    user.Email,
		Password: user.Password,
	}
}

type VerifyCode struct {
	mgm.DefaultModel `bson:"inline"`
	Code             int32  `json:"code" bson:"code"`
	Email            string `json:"email" bson:"email"`
}

func NewVerifyCode(code *VerifyCode) *VerifyCode {
	return &VerifyCode{
		Code:  code.Code,
		Email: code.Email,
	}
}
