package requests

type DeleteUserReq struct {
	IDs []int64 `json:"ids" validate:"required"`
}
