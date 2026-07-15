package composers

import (
	"gocas/modules/admins/users_group/transport/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
)

type composerUserGroupHdl interface {
	ListUserGroupHdl() fiber.Handler
}

func ComposerUserGroupHdlService(serviceCtx gosctx.ServiceContext) composerUserGroupHdl {
	// db := serviceCtx.MustGet(configs.KeyCompGorm).(gormc.GormComponent).GetDB()

	// repo := repository.NewUserGroupRepo(db)

	// usc := usecase.NewUserGroupUseCase(repo)

	hdl := handlers.NewUserGroupHdl()

	return hdl
}
