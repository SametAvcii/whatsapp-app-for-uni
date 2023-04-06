package repository

import (
	models "whatsapp-app/model"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IFacultyRepository interface {
	Create(NewFaculty *models.Faculty) error
	FindByID(id uint16) (*models.Faculty, error)
	IsExist(name string) bool
}

type FacultyRepository struct {
	db *mongo.Database
}

func NewFacultyRepository(db *mongo.Database) *FacultyRepository {
	return &FacultyRepository{db: db}
}

func (f *FacultyRepository) Create(NewFaculty *models.Faculty) error {
	faculty := models.NewFaculty(NewFaculty)
	err := mgm.Coll(faculty).Create(faculty)
	return err
}

func (f *FacultyRepository) FindByID(code uint16) (*models.Faculty, error) {
	faculty := &models.Faculty{}
	err := mgm.Coll(faculty).First(bson.M{"code": code}, faculty)

	return faculty, err
}

func (f *FacultyRepository) IsExist(code uint16) bool {
	faculty := &models.Faculty{}
	err := mgm.Coll(faculty).First(bson.M{"code": code}, faculty)
	if err != nil {
		return false
	}

	return true
}
