package usecase

import (
	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users_group/entity"
)

type repo interface {
	GetUserGroupByIdRepo(id int64) (*entity.UserGroup, error)
}

type userGroupUsc struct {
	repo repo
}

func NewUserGroupUseCase(repo repo) *userGroupUsc {
	return &userGroupUsc{repo: repo}
}

func (u *userGroupUsc) ListUserGroupUsc() {
}

func (u *userGroupUsc) FindOneUserGroupUsc(id int) (*entity.UserGroup, error) {
	data, err := u.repo.GetUserGroupByIdRepo(int64(id))

	if err != nil {
		if err == errs.ErrInternalServer {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}
