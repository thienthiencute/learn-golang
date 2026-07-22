package composers

import (
	"gocas/modules/admins/users_group/repository"
	"gocas/modules/admins/users_group/transport/api"
	"gocas/modules/admins/users_group/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/component/gormc"
	"github.com/teoit/gosctx/configs"
)
// composerUserGroupApi: interface định nghĩa các API handler mà module user group cung cấp
// (dùng interface để route chỉ phụ thuộc vào contract, không phụ thuộc implementation cụ thể)
type composerUserGroupApi interface {
	ListUserGroupApi() fiber.Handler
	CreateUserGroupApi() fiber.Handler
	UpdateUserGroupApi() fiber.Handler
	UpdateStatusUserGroupApi() fiber.Handler
	DetailUserGroupApi() fiber.Handler
	DeleteUserGroupApi() fiber.Handler
}

// ComposerUserGroupApiService: khởi tạo toàn bộ chain dependency cho module user group API
// theo hướng Clean Architecture: Repo -> Usecase -> Api (handler)
func ComposerUserGroupApiService(serviceCtx gosctx.ServiceContext) composerUserGroupApi {

	// Lấy DB connection (GORM) đã được đăng ký sẵn trong serviceCtx
	db := serviceCtx.MustGet(configs.KeyCompGorm).(gormc.GormComponent).GetDB()

	// Tầng Repository: chịu trách nhiệm thao tác trực tiếp với DB
	repo := repository.NewUserGroupRepo(db)

	// Tầng Usecase: chứa business logic, gọi xuống repo
	usc := usecase.NewUserGroupApiUsc(repo)

	// Tầng Api : nhận request, gọi usecase, trả response
	api := api.NewUserGroupApi(usc)

	return api
}
