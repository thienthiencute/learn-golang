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

type composerUserGroupApi interface {
	ListUserGroupApi() fiber.Handler
	CreateUserGroupApi() fiber.Handler
	UpdateUserGroupApi() fiber.Handler
	DetailUserGroupApi() fiber.Handler
}

func ComposerUserGroupApiService(serviceCtx gosctx.ServiceContext) composerUserGroupApi {
	db := serviceCtx.MustGet(configs.KeyCompGorm).(gormc.GormComponent).GetDB()

	repo := repository.NewUserGroupRepo(db)
	usc := usecase.NewUserGroupApiUsc(repo)

	api := api.NewUserGroupApi(usc)

	return api
}
