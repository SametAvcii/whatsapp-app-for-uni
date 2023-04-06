package middlewares

import (
	"whatsapp-app/internal/config/db"
	"whatsapp-app/internal/repository"
	"whatsapp-app/internal/utils/response"
	models "whatsapp-app/model"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenUser, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(401, response.Response(401, "Oturum bulunamadı."))
		}

		claims, ok := tokenUser.Claims.(*models.JWTCustomClaims)
		if !ok {
			return c.JSON(401, response.Response(401, "Oturum bulunamadı."))
		}

		db := db.Connect()

		verifiedUser, err := repository.NewUserRepository(db).FindWithApiKey(claims.ApiKey)
		if err != nil {
			return c.JSON(401, response.Response(401, "Kullanıcı bulunamadı."))
		}

		c.Set("verifiedUser", verifiedUser)
		return next(c)
	}
}
