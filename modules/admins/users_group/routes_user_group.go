package users_group

import (
	"gocas/modules/admins/users_group/composers"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
)

// SetupRoutesUserGroup: đăng ký các route cho module (user group / roles)
func SetupRoutesUserGroup(app *fiber.App, serviceCtx gosctx.ServiceContext) {
	// Nhóm route cho giao diện admin (render HTML/view)
	groupHdl := app.Group("/admins/roles")
	{
		comp := composers.ComposerUserGroupHdlService(serviceCtx)
		groupHdl.Get("/list", comp.ListUserGroupHdl()).Name("ecommerce.users_group.list")
		groupHdl.Get("/update/:id", comp.UpdateUserGroupHdl()).Name("ecommerce.users_group.update")
	}

	// Nhóm route cho API (trả JSON)
	groupApi := app.Group("/api/admins/roles")
	{
		comp := composers.ComposerUserGroupApiService(serviceCtx)
		groupApi.Post("/list", comp.ListUserGroupApi()).Name("ecommerce.users_group.api.list")
		groupApi.Post("/create", comp.CreateUserGroupApi()).Name("ecommerce.users_group.api.create")
		groupApi.Post("/update", comp.UpdateUserGroupApi()).Name("ecommerce.users_group.api.update")
		groupApi.Patch("/update-status", comp.UpdateStatusUserGroupApi()).Name("ecommerce.users_group.api.update-status")
		groupApi.Delete("/delete", comp.DeleteUserGroupApi()).Name("ecommerce.users_group.api.delete")
	}
}





