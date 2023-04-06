package db

import (
	"time"
	configs "whatsapp-app/internal/config"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	_ = mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 12 * time.Second}, configs.DbName, options.Client().
		ApplyURI("mongodb+srv://"+configs.DbUser+":"+configs.DbPass+"@whatsapp-app.hecfb4r.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPIOptions))

	client, err := mgm.NewClient(options.Client().ApplyURI("mongodb+srv://" + configs.DbUser + ":" + configs.DbPass + "@whatsapp-app.hecfb4r.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPIOptions))

	if err != nil {
		panic(err)
	}

	db := client.Database(configs.DbName)
	return db
}
