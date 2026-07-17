package usecase

import (
	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users_group/entity"
	"gocas/modules/admins/users_group/mapping"
	"gocas/modules/admins/users_group/transport/responses"
)

type repoApi interface {
	ListUserGroupRepo() (*[]entity.UserGroup, error)
	CreateUserGroupRepo(data *entity.UserGroup) error
	UpdateUserGroupRepo(id int64, data map[string]interface{}) error
	GetUserGroupByIdRepo(id int64) (*entity.UserGroup, error)
}

type userGroupApiUsc struct {
	repoApi repoApi
}

func NewUserGroupApiUsc(repoApi repoApi) *userGroupApiUsc {
	return &userGroupApiUsc{repoApi: repoApi}
}

func (u *userGroupApiUsc) ListApiUserGroupUsc() (*[]responses.UserGroupResp, error) {
	data, err := u.repoApi.ListUserGroupRepo()
	if err != nil {
		if err == errs.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	result := mapping.MapperListUserGroup(data)
	
	return result, nil

}

func (u *userGroupApiUsc) CreateApiUserGroupUsc(name, description string, status int) error {
	userGroup := &entity.UserGroup{
		Name:        name,
		Description: description,
		Status:      status,
	}

	err := u.repoApi.CreateUserGroupRepo(userGroup)
	if err != nil {
		return err
	}

	return nil
}

func (u *userGroupApiUsc) UpdateApiUserGroupUsc(id int64, name, description string, status int) error {
	data := map[string]interface{}{
		"name":        name,
		"description": description,
		"status":      status,
	}

	err := u.repoApi.UpdateUserGroupRepo(id, data)
	if err != nil {
		return err
	}

	return nil
}

func (u *userGroupApiUsc) DetailApiUserGroupUsc(id int64) (*entity.UserGroup, error) {
	data, err := u.repoApi.GetUserGroupByIdRepo(id)
	if err != nil {
		if err == errs.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}
