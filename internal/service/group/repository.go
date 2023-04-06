package group

import "whatsapp-app/internal/repository"

type Repository struct {
	Faculty    *repository.FacultyRepository
	Department *repository.DepartmentRepository
	Group      *repository.GroupRepository
}
