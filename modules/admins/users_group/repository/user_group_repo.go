package repository

import (
	"context"
	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users_group/entity"
	"gocas/pkg/golangviet/generics"
	"gocas/pkg/golangviet/utils"

	"gorm.io/gorm"
)

// userGroupRepo: struct implement tầng Repository, chịu trách nhiệm thao tác trực tiếp với DB
// thông qua GORM. Đây là tầng thấp nhất trong kiến trúc (Repo -> Usecase -> Api/Handler)
type userGroupRepo struct {
	db *gorm.DB
}

// NewUserGroupRepo: constructor - nhận DB connection từ ngoài vào (Dependency Injection)
func NewUserGroupRepo(db *gorm.DB) *userGroupRepo {
	return &userGroupRepo{db: db}
}

// ListUserGroupRepo: lấy toàn bộ danh sách nhóm quyền
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

// find user conditions...
func (s *userGroupRepo) FindUserGroupRepo(ctx context.Context, filter *utils.Filters) ([]*entity.UserGroup, error) {
	tableName := entity.UserGroup{}.TableName()
	return generics.FindGeneric[entity.UserGroup](ctx, s.db, tableName, filter)
}

// CreateUserGroupRepo: tạo mới 1 nhóm quyền, data được insert trực tiếp vào DB
func (s *userGroupRepo) CreateUserGroupRepo(data *entity.UserGroup) error {
	result := s.db.Table(entity.UserGroup{}.TableName()).Create(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateUserGroupRepo: cập nhật theo id, data là map nên chỉ update đúng field được truyền vào
// (Updates(map) trong GORM chỉ update field có trong map, khác với Updates(struct) sẽ bỏ qua
func (s *userGroupRepo) UpdateUserGroupRepo(id int64, data map[string]interface{}) error {
	result := s.db.Model(&entity.UserGroup{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserGroupByIdRepo: lấy chi tiết 1 nhóm quyền theo id
// Dùng First() -> tự động trả gorm.ErrRecordNotFound nếu không tìm thấy, được convert
// sang lỗi custom errs.ErrRecordNotFound để đồng nhất lỗi trong toàn hệ thống
// (Usecase/Api không cần biết đến gorm.ErrRecordNotFound, chỉ cần biết errs.ErrRecordNotFound)
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

// DeleteUserGroupRepo: xóa nhiều nhóm quyền cùng lúc bằng 1 câu SQL (WHERE id IN (...))
// Unscoped(): hard delete, xóa thật khỏi DB (bỏ qua soft delete nếu entity có DeletedAt)
func (s *userGroupRepo) DeleteUserGroupRepo(ids []int64) error {
	result := s.db.Table(entity.UserGroup{}.TableName()).
		Unscoped().
		Where("id IN ?", ids).
		Delete(&entity.UserGroup{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
