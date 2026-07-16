package requests

type CreateUserGroupReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateUserGroupReq struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}