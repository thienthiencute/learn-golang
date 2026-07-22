package repository

import (
	// "context"

	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users/entity"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}
}

func (s *userRepo) ListUserRepo() (*[]entity.User, error) {
	var data []entity.User
	result := s.db.Table(entity.User{}.TableName()).Preload("UserCreated").Preload("Role").Order("id desc").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (s *userRepo) ListRoles() ([]entity.Role, error) {
	var roles []entity.Role
	err := s.db.Table(entity.Role{}.TableName()).Find(&roles).Error
	return roles, err
}

// detail
func (s *userRepo) DetailUserRepo(id int64) (*entity.User, error) {
	var data entity.User
	result := s.db.Table(entity.User{}.TableName()).Preload("Role").Where("id = ?", id).First(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (s *userRepo) UpdateUserRepo(id int64, data map[string]interface{}) error {
	result := s.db.Model(&entity.User{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errs.ErrRecordNotFound
	}
	return nil
}

func (s *userRepo) DeleteUserRepo(ids []int64) error {
	result := s.db.Table(entity.User{}.TableName()).Where("id IN ?", ids).Delete(&entity.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// create
func (s *userRepo) InsertUserRepo(data *entity.User) (*int64, error) {
	result := s.db.Table(entity.User{}.TableName()).Create(data)
	if result.Error != nil {
		return nil, result.Error

	}

	return &data.ID, nil
}

// func (r *userRepo) FindUserRepo(ctx context.Context, filter *){

// }
