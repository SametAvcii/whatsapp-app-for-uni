package main

import (
	"whatsapp-app/internal/config/db"
	"whatsapp-app/internal/config/redis"
	"whatsapp-app/router"

	"github.com/labstack/echo/v4"
)

func init() {
	db.Connect()
}

func main() {

	db := db.Connect()
	e := echo.New()
	client := redis.NewClient()
	router.Init(e, db, client)

	e.Logger.Fatal(e.Start(":8080"))

}
