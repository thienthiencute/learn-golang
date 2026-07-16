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

func (s *userGroupRepo) CreateUserGroupRepo(data *entity.UserGroup) error {
	result := s.db.Table(entity.UserGroup{}.TableName()).Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *userGroupRepo) UpdateUserGroupRepo(id int64, data map[string]interface{}) error {
	result := s.db.Table(entity.UserGroup{}.TableName()).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *userGroupRepo) GetUserGroupByIdRepo(id int64) (*entity.UserGroup, error) {
	var data entity.UserGroup
	result := s.db.Table(entity.UserGroup{}.TableName()).Where("id = ?", id).First(&data)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errs.ErrRecordNotFound
		}
		return nil, result.Error
	}
	return &data, nil
}
