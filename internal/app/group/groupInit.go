package group

import (
	"whatsapp-app/internal/repository"
	"whatsapp-app/internal/service"
	"whatsapp-app/internal/service/group"
	groupService "whatsapp-app/internal/service/group"
	"whatsapp-app/internal/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

func GroupInit(db *mongo.Database) IGroupHandler {
	utils := utils.NewUtils()
	groupRepository := repository.NewGroupRepository(db)
	departmentRepository := repository.NewDepartmentRepository(db)
	facultyRepository := repository.NewFacultyRepository(db)

	repository := group.Repository{Faculty: facultyRepository, Department: departmentRepository, Group: groupRepository}

	service := groupService.NewGroupService(repository, service.Service{Utils: utils})
	handler := NewGroupHandler(service, utils)
	return handler
}
