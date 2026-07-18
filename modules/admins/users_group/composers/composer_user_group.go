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

type composerUserGroupHdl interface {
	ListUserGroupHdl() fiber.Handler
	UpdateUserGroupHdl() fiber.Handler
}

func ComposerUserGroupHdlService(serviceCtx gosctx.ServiceContext) composerUserGroupHdl {
	db := serviceCtx.MustGet(configs.KeyCompGorm).(gormc.GormComponent).GetDB()

	repo := repository.NewUserGroupRepo(db)

	usc := usecase.NewUserGroupUseCase(repo)

	hdl := handlers.NewUserGroupHdl(usc)

	return hdl
}
