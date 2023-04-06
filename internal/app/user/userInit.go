package user

import (
	"whatsapp-app/cache"
	"whatsapp-app/internal/repository"
	"whatsapp-app/internal/service"
	UserService "whatsapp-app/internal/service/user"
	utils "whatsapp-app/internal/utils"

	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserInit(db *mongo.Database, client *redis.Client) IUserHandler {
	utils := utils.NewUtils()
	repository := repository.NewUserRepository(db)
	cache := cache.NewCache(client)
	service := UserService.NewUserService(repository, service.Service{Utils: utils}, cache)
	handler := NewUserHandler(service, utils)
	return handler
}
