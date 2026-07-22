package users

import (
	"gocas/modules/admins/users/composers"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
)

// views -> root -> server -> server.go (trong internal) -> routes_user.go -> handler (show ra giao dien) !! api(khong mo ra xem duoc) -> usecase -> repository -> entity

func SetupRoutesUser(app *fiber.App, serviceCtx gosctx.ServiceContext) {
	group := app.Group("/admins/users")
	{
		comp := composers.ComposerUserService(serviceCtx)
		group.Get("/list", comp.ListUserHdl()).Name("ecommerce.users.list")
		group.Get("/update/:id", comp.UpdateUserHdl()).Name("ecommerce.users.update")
	}

	groupApi := app.Group("/api/admins/users")
	{
		compApi := composers.ComposerUserApi(serviceCtx)
		groupApi.Post("/list", compApi.ListUserApi()).Name("ecommerce.users.api.list")
		groupApi.Post("/create", compApi.CreateUserApi()).Name("ecommerce.users.api.create")
		groupApi.Post("/update", compApi.UpdateUserApi()).Name("ecommerce.users.api.update")
		groupApi.Patch("/update-status", compApi.UpdateStatusUserApi()).Name("ecommerce.users.api.update-status")
		groupApi.Delete("/delete", compApi.DeleteUserApi()).Name("ecommerce.users.api.delete")
	}

}
