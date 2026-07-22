package composers

import (
	"gocas/modules/admins/users_group/repository"
	"gocas/modules/admins/users_group/transport/handlers"
	"gocas/modules/admins/users_group/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/component/gormc"
	"github.com/teoit/gosctx/configs"
)

// composerUserGroupHdl: interface định nghĩa các handler render VIEW (HTML) cho module user group
// (khác với composerUserGroupApi ở trên - đây là cho giao diện admin, không trả JSON)
type composerUserGroupHdl interface {
	ListUserGroupHdl() fiber.Handler
	UpdateUserGroupHdl() fiber.Handler
}

// ComposerUserGroupHdlService: wiring dependency cho module user group phía VIEW
// Luồng: DB -> Repo -> Usecase -> Handler (Hdl)
func ComposerUserGroupHdlService(serviceCtx gosctx.ServiceContext) composerUserGroupHdl {
	// Lấy DB connection từ DI container
	db := serviceCtx.MustGet(configs.KeyCompGorm).(gormc.GormComponent).GetDB()

	// Tầng Repository: thao tác DB (tái sử dụng lại repo giống bên Api)
	repo := repository.NewUserGroupRepo(db)

	// Tầng Usecase: business logic riêng cho phần view/admin
	usc := usecase.NewUserGroupUseCase(repo)

	// Tầng Handler: nhận request, gọi usecase, render ra HTML/view
	hdl := handlers.NewUserGroupHdl(usc)

	return hdl
}
