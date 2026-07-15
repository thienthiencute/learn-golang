package usecase

import "gocas/modules/admins/users/entity"

type UserRepo interface {
	InsertUserRepo(data *entity.User) (*int64, error)
}

// type RoleRope interface{

// }
type userUseCase struct {
	userRepo UserRepo
	// roleRope RoleRope
}

// func NewUserUseCase(userRepo UserRepo,roleRepo RoleRope) *userUseCase {
// 	return &userUseCase{userRepo: userRepo, roleRope: roleRepo}

// }

func NewUserUseCase(userRepo UserRepo) *userUseCase {
	return &userUseCase{userRepo: userRepo}

}