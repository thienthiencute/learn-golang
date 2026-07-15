package users_group

import (
	"gocas/modules/admins/users_group/composers"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
)

func SetupRoutesUserGroup(app *fiber.App, serviceCtx gosctx.ServiceContext) {
	groupHdl := app.Group("/admins/users-group")
	{
		comp := composers.ComposerUserGroupHdlService(serviceCtx)
		groupHdl.Get("/list", comp.ListUserGroupHdl()).Name("ecommerce.users_group.list")
	}

	groupApi := app.Group("/api/admins/users-group")
	{
		comp := composers.ComposerUserGroupApiService(serviceCtx)
		groupApi.Post("/list", comp.ListUserGroupApi()).Name("ecommerce.users_group.api.list")
	}
}

//