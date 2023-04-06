package request

type UserRegisterDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginDTO struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserVerifyDTO struct {
	Email string `json:"email" validate:"required"`
}
type UserVerifyEmailDTO struct {
	Email string `json:"email" validate:"required"`
	Code  string `json:"code" validate:"required"`
}
