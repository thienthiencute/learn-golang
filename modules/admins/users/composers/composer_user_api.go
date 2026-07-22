package composers

import (
	"gocas/modules/admins/users/repository"
	"gocas/modules/admins/users/transport/api"
	"gocas/modules/admins/users/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/component/gormc"
	"github.com/teoit/gosctx/configs"
)

type composerUserApi interface {
	ListUserApi() fiber.Handler
	CreateUserApi() fiber.Handler
	UpdateStatusUserApi() fiber.Handler
	UpdateUserApi() fiber.Handler
	DeleteUserApi() fiber.Handler
}

func ComposerUserApi(serviceCtx gosctx.ServiceContext) composerUserApi {
	db := serviceCtx.MustGet(configs.KeyCompGorm).(gormc.GormComponent).GetDB()

	repo := repository.NewUserRepo(db)

	usc := usecase.NewUserApiUsc(repo)

	hdl := api.NewUserApi(usc)

	return hdl
}
