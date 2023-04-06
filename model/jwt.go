package models

import "github.com/golang-jwt/jwt"

type JWTCustomClaims struct {
	ApiKey string `json:"api_key"`
	jwt.StandardClaims
}
