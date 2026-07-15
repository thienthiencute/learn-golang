package handlers

import (
	"gocas/pkg/golangviet/templates"
	"gocas/views/admins/users_group"

	"github.com/gofiber/fiber/v2"
)



type userGroupHdl struct {
}

func NewUserGroupHdl() *userGroupHdl {
	return &userGroupHdl{}
}

func (h *userGroupHdl) ListUserGroupHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return templates.Render(c, users_group.Index())
	}
}
