package handlers

import (
	"gocas/pkg/golangviet/templates"
	"gocas/views/admins/users"

	"github.com/gofiber/fiber/v2"
)

// trong day chi genre ra html, con api la xu ly

type UserUsc interface{}

type userHdl struct {
	userUsc UserUsc
}

func NewUserHdl(userUsc UserUsc) *userHdl {
	return &userHdl{userUsc: userUsc}

}

func (h *userHdl) ListUserHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return templates.Render(c, users.Index())
		// return c.SendString("Hi")
	}
}
