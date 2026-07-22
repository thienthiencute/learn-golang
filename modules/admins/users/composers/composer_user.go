package composers

import (
	"gocas/modules/admins/users/repository"
	"gocas/modules/admins/users/transport/handlers"
	"gocas/modules/admins/users/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/component/gormc"
	"github.com/teoit/gosctx/configs"
)

type composerUserHdl interface {
	ListUserHdl() fiber.Handler
	UpdateUserHdl() fiber.Handler
}

func ComposerUserService(serviceCtx gosctx.ServiceContext) composerUserHdl {
	db := serviceCtx.MustGet(configs.KeyCompGorm).(gormc.GormComponent).GetDB()

	repo := repository.NewUserRepo(db)

	usc := usecase.NewUserUseCase(repo)

	hdl := handlers.NewUserHdl(usc)

	return hdl
}
