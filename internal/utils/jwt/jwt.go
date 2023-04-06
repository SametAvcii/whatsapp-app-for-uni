package jwt

import (
	configs "whatsapp-app/internal/config"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var JWTConfig = middleware.JWTConfig{
	Claims:     &JwtCustomClaims{},
	SigningKey: []byte(configs.TokenSecret),
}

type JwtCustomClaims struct {
	ID primitive.ObjectID `json:"id"`
	jwt.StandardClaims
}
