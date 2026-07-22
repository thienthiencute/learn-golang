package usecase

import "gocas/modules/admins/users/entity"

type UserRepo interface {
	InsertUserRepo(data *entity.User) (*int64, error)
	DetailUserRepo(id int64) (*entity.User, error)
	ListRoles() ([]entity.Role, error)
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

func (s *userUseCase) DetailUserUsc(id int64) (*entity.User, error) {
	return s.userRepo.DetailUserRepo(id)
}

func (s *userUseCase) ListRolesUsc() ([]entity.Role, error) {
	return s.userRepo.ListRoles()
}