package usecase

import (
	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users_group/entity"
	"gocas/modules/admins/users_group/mapping"
	"gocas/modules/admins/users_group/transport/responses"
)

type repoApi interface {
	ListUserGroupRepo() (*[]entity.UserGroup, error)
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
