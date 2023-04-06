package response

type NewDepartmentDTO struct {
	FacultyName string `bson:"faculty_name" json:"faculty_name"`
	Name        string `bson:"department_name" json:"department_name"`
	Code        uint16 `json:"code"`
}
