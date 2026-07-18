package handlers

import (
	"gocas/modules/admins/users_group/entity"
	"gocas/pkg/golangviet/templates"
	"gocas/views/admins/users_group"
	"gocas/views/pages"

	"github.com/gofiber/fiber/v2"
)

type userGroupUsc interface {
	FindOneUserGroupUsc(id int) (*entity.UserGroup, error)
}

type userGroupHdl struct {
	userGroupUsc userGroupUsc
}

func NewUserGroupHdl(userGroupUsc userGroupUsc) *userGroupHdl {
	return &userGroupHdl{userGroupUsc: userGroupUsc}
}

func (h *userGroupHdl) ListUserGroupHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return templates.Render(c, users_group.Index())
	}
}

func (h *userGroupHdl) UpdateUserGroupHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil || id <= 0 {
			return templates.Render(c, errorpages.NotFound404("/admins/users-group/list"))
		}

		detailUserGroup, err := h.userGroupUsc.FindOneUserGroupUsc(id)

		if err != nil {
			return templates.Render(c, errorpages.NotFound404("/admins/users-group/list"))
		}

		return templates.Render(c, users_group.Update(detailUserGroup))
	}
}
