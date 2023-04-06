package response

import (
	models "whatsapp-app/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRegisterDTO struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	SchoolID int32  `json:"school_id"`
	Email    string `json:"email"`
}

type UserLoginDTO struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	SchoolID int32              `json:"school_id"`
	Email    string             `json:"email"`
	Verified bool               `json:"verified"`
	Token    string             `json:"token"`
}

func (u *UserLoginDTO) Convert(user *models.User, token string) {
	u.ID = user.ID
	u.Name = user.Name
	u.Email = user.Email
	u.SchoolID = user.SchoolID
	u.Verified = user.Verified
	u.Token = token

}

type UserVerifyDTO struct {
	Message string `json:"message"`
}
