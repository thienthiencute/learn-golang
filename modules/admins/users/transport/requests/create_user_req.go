package requests

type CreateUserReq struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	RoleID    int64  `json:"role_id" validate:"required"`
	Status    int    `json:"status"`
}
