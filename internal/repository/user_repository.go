package repository

import (
	"errors"
	"time"
	models "whatsapp-app/model"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collectionName string = "basic_auth"

	mongoQueryTimeout = 10 * time.Second
)

type IUserRepository interface {
	FindAllUser() ([]models.User, error)
	IsDuplicateSchoolID(id int32) bool
	CreateUser(newUser *models.User) error
	Login(school_id int32) (*models.User, error)
	FindUserWithEmail(email string) (*models.User, error)
	IsExistWithEmail(email string) bool
	FindWithApiKey(apiKey string) (*models.User, error)
	UpdateUser(user *models.User) error
}

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) CreateUser(newUser *models.User) error {
	user := models.NewUser(newUser)
	err := mgm.Coll(user).Create(user)
	return err
}

func (u *UserRepository) FindAllUser() ([]models.User, error) {

	var users []models.User
	user := &models.User{}
	err := mgm.Coll(user).SimpleFind(&users, bson.M{})
	return users, err
}

func (u *UserRepository) IsDuplicateSchoolID(id int32) bool {
	user := &models.User{}
	err := mgm.Coll(user).First(bson.M{"school_id": id}, user)
	if err != nil {
		return false
	}
	return true
}

func (u *UserRepository) IsExistWithEmail(email string) bool {
	user := &models.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)
	if err != nil {
		return false
	}
	return true
}

func (u *UserRepository) FindUserWithEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := mgm.Coll(user).First(bson.M{"email": email}, user)

	return user, err
}

func (u *UserRepository) Login(school_id int32) (*models.User, error) {

	user := &models.User{}
	err := mgm.Coll(user).First(bson.M{"school_id": school_id}, user)
	if err != nil {
		return user, errors.New("Kullanıcı bulunamadı.")
	}
	return user, nil
}

func (u *UserRepository) UpdateUser(user *models.User) error {
	err := mgm.Coll(user).Update(user)
	return err
}

func (u *UserRepository) FindWithApiKey(apiKey string) (*models.User, error) {

	user := &models.User{}
	err := mgm.Coll(user).First(bson.M{"api_key": apiKey}, user)
	return user, err

}
