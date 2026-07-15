package repository

import (
	// "context"

	"gocas/modules/admins/users/entity"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db: db}

}

//create
func (s *userRepo) InsertUserRepo(data *entity.User) (*int64, error) {
	result := s.db.Table(entity.User{}.TableName()).Create(data)
	if result.Error != nil {
		return nil, result.Error

	}

	return &data.ID, nil
}

// func (r *userRepo) FindUserRepo(ctx context.Context, filter *){

// }
