package usecase

type repo interface {
}

type userGroupUsc struct {
	repo repo
}

func NewUserGroupUseCase(repo repo) *userGroupUsc {
	return &userGroupUsc{repo: repo}
}

func (u *userGroupUsc) ListUserGroupUsc() {
}
