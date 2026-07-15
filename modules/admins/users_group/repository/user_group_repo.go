package repository

import (
	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users_group/entity"

	"gorm.io/gorm"
)

type userGroupRepo struct {
	db *gorm.DB
}

func NewUserGroupRepo(db *gorm.DB) *userGroupRepo {
	return &userGroupRepo{db: db}
}

func (s *userGroupRepo) ListUserGroupRepo() (*[]entity.UserGroup, error) {
	var data []entity.UserGroup
	result := s.db.Table(entity.UserGroup{}.TableName()).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errs.ErrRecordNotFound
	}

	return &data, nil
}
