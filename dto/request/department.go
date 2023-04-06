package request

type NewDepartmentDTO struct {
	Name        string `json:"name" validate:"required"`
	FacultyCode uint16 `json:"faculty_code" validate:"required"`
	Code        uint16 `json:"code" validate:"required"`
}
