package internal

import (
	"fmt"
	"gocas/conf"
	"gocas/modules/admins/users"
	"gocas/modules/admins/users_group"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx"
)

func RoutesServer(app *fiber.App, serviceCtx gosctx.ServiceContext) {
	users.SetupRoutesUser(app, serviceCtx)
	users_group.SetupRoutesUserGroup(app, serviceCtx)

	app.Static("/static", fmt.Sprintf("./%s", conf.UploadPathPublic))
}
