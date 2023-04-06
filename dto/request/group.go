package request

type NewGroupDTO struct {
	Name           string `json:"name" validate:"required"`
	DepartmentCode uint16 `json:"department_code" validate:"required"`
	Link           string `json:"link" validate:"required"`
}
type VerifyGroup struct {
	ID       uint16 `json:"id" validate:"required"`
	IsVerify bool  `json:"is_verify" validate:"required"`
}
