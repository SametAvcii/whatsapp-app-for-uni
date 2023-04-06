package repository

import (
	models "whatsapp-app/model"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IDepartmentRepository interface {
	FindByID(id uint16) (*models.Department, error)
	IsExist(name string) bool
}

type DepartmentRepository struct {
	db *mongo.Database
}

func NewDepartmentRepository(db *mongo.Database) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (d *DepartmentRepository) FindByID(code uint16) (*models.Department, error) {
	department := &models.Department{}
	err := mgm.Coll(department).First(bson.M{"code": code}, department)

	return department, err
}
func (d *DepartmentRepository) IsExist(code uint16) bool {
	department := &models.Department{}
	err := mgm.Coll(department).First(bson.M{"code": code}, department)
	if err != nil {
		return false
	}

	return true
}
func (f *DepartmentRepository) Create(newDepartment *models.Department) error {
	department := models.NewDepartment(newDepartment)
	err := mgm.Coll(department).Create(department)
	return err
}
