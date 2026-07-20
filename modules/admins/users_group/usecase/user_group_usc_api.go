package usecase

import (
	"context"
	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users_group/entity"
	"gocas/modules/admins/users_group/mapping"
	"gocas/modules/admins/users_group/transport/requests"
	"gocas/modules/admins/users_group/transport/responses"
)

type userGroupApi interface {
	ListUserGroupRepo() (*[]entity.UserGroup, error)
	CreateUserGroupRepo(data *entity.UserGroup) error
	UpdateUserGroupRepo(id int64, data map[string]interface{}) error
	GetUserGroupByIdRepo(id int64) (*entity.UserGroup, error)
	DeleteUserGroupRepo(id int64) error
}

type userGroupApiUsc struct {
	userGroupApi userGroupApi
}

func NewUserGroupApiUsc(userGroupApi userGroupApi) *userGroupApiUsc {
	return &userGroupApiUsc{userGroupApi: userGroupApi}
}

func (u *userGroupApiUsc) ListApiUserGroupUsc() (*[]responses.UserGroupResp, error) {
	data, err := u.userGroupApi.ListUserGroupRepo()
	if err != nil {
		if err == errs.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	result := mapping.MapperListUserGroup(data)

	return result, nil

}

func (u *userGroupApiUsc) CreateApiUserGroupUsc(ctx context.Context, req *requests.UserGroupCreation) error {
	userGroup := &entity.UserGroup{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}

	err := u.userGroupApi.CreateUserGroupRepo(userGroup)

	if err != nil {
		return err
	}

	return nil
}

func (u *userGroupApiUsc) UpdateApiUserGroupUsc(ctx context.Context, req *requests.UserGroupUpdateReq) error {
	data := map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
		"status":      req.Status,
	}

	err := u.userGroupApi.UpdateUserGroupRepo(int64(req.ID), data)
	if err != nil {
		return err
	}

	return nil
}

func (u *userGroupApiUsc) DetailApiUserGroupUsc(ctx context.Context, req *requests.UserGroupDetailReq) (*entity.UserGroup, error) {
	data, err := u.userGroupApi.GetUserGroupByIdRepo(req.ID)
	if err != nil {
		if err == errs.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}

func (u *userGroupApiUsc) DeleteApiUserGroupUsc(ctx context.Context, id int64) error {
	err := u.userGroupApi.DeleteUserGroupRepo(id)
	if err != nil {
		return err
	}
	return nil
}
