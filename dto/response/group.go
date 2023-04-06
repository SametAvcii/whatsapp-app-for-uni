package response

type NewGroupDTO struct {
	DepartmentName string `json:"department_name"`
	Name           string `json:"group_name"`
	Link           string `json:"link"`
}

type GroupDTO struct {
	ID         string `json:"id"`
	Name       string `json:"group_name"`
	Link       string `json:"link"`
	IsVerified bool   `json:"is_verified"`
}
type GetGroupsDTO struct {
	Groups []GroupDTO
}
