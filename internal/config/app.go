package configs

import (
	models "whatsapp-app/model"

	"github.com/labstack/echo/v4/middleware"
)

const DbName = "whatsapp-app"
const DbPass = "T.C.1923a"
const DbUser = "root"
const TokenSecret = "group-app-firat"
const EmailAddress = "wpapp23@outlook.com"
const EmailPass = "Appsecret1."

var JWTConfigApp = middleware.JWTConfig{
	Claims:     &models.JWTCustomClaims{},
	SigningKey: []byte(TokenSecret),
}
