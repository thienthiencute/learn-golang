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
	}

}
