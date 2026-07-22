package usecase

import (
	"context"
	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users/entity"
	"gocas/modules/admins/users/mapping"
	"gocas/modules/admins/users/transport/requests"
	"gocas/modules/admins/users/transport/responses"
)

type userApiRepo interface {
	ListUserRepo() (*[]entity.User, error)
	InsertUserRepo(data *entity.User) (*int64, error)
	UpdateUserRepo(id int64, data map[string]interface{}) error
	DeleteUserRepo(ids []int64) error
}

type UserApiUsc interface {
	ListApiUserUsc(ctx context.Context) (*[]responses.UserResp, error)
	CreateApiUserUsc(ctx context.Context, data *entity.User) (*int64, error)
	UpdateStatusApiUserUsc(ctx context.Context, req *requests.UserUpdateStatusReq) error
	UpdateApiUserUsc(ctx context.Context, req *requests.UpdateUserReq) error
	DeleteApiUserUsc(ctx context.Context, req *requests.DeleteUserReq) error
}

type userApiUsc struct {
	userApiRepo userApiRepo
}

func NewUserApiUsc(userApiRepo userApiRepo) *userApiUsc {
	return &userApiUsc{userApiRepo: userApiRepo}
}

func (u *userApiUsc) ListApiUserUsc(ctx context.Context) (*[]responses.UserResp, error) {
	data, err := u.userApiRepo.ListUserRepo()
	if err != nil {
		if err == errs.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	result := mapping.MapperListUser(data)

	return result, nil
}

func (u *userApiUsc) CreateApiUserUsc(ctx context.Context, data *entity.User) (*int64, error) {
	return u.userApiRepo.InsertUserRepo(data)
}

func (u *userApiUsc) UpdateStatusApiUserUsc(ctx context.Context, req *requests.UserUpdateStatusReq) error {
	data := map[string]interface{}{
		"status": req.Status,
	}

	err := u.userApiRepo.UpdateUserRepo(int64(req.ID), data)
	if err != nil {
		return err
	}

	return nil
}

func (u *userApiUsc) UpdateApiUserUsc(ctx context.Context, req *requests.UpdateUserReq) error {
	data := map[string]interface{}{
		"email":      req.Email,
		"first_name": req.FirstName,
		"last_name":  req.LastName,
		"status":     req.Status,
	}

	err := u.userApiRepo.UpdateUserRepo(int64(req.ID), data)
	if err != nil {
		return err
	}

	return nil
}

func (u *userApiUsc) DeleteApiUserUsc(ctx context.Context, req *requests.DeleteUserReq) error {
	return u.userApiRepo.DeleteUserRepo(req.IDs)
}
