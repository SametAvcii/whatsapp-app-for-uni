package router

import (
	"whatsapp-app/internal/app/group"
	"whatsapp-app/internal/app/user"
	configs "whatsapp-app/internal/config"
	"whatsapp-app/internal/middlewares"

	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go.mongodb.org/mongo-driver/mongo"
)

func Init(router *echo.Echo, tx *mongo.Database, client *redis.Client) {
	userHandler := user.UserInit(tx, client)

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType},
	}))

	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	//email.POST("/send", userHandler.SendEmail)
	router.POST("/verify", userHandler.VerifyEmail)

	auth := router.Group("")

	auth.Use(middleware.JWTWithConfig(configs.JWTConfigApp))
	auth.Use(middlewares.VerifyToken)
	groupHandler := group.GroupInit(tx)
	auth.POST("/group", groupHandler.NewGroup)
	auth.POST("/department", groupHandler.NewDepartment)
	auth.POST("/faculty", groupHandler.NewFaculty)
	auth.GET("/groups", groupHandler.GetGroups)
	auth.PUT("/verify/group", groupHandler.VerifyGroup)
}
