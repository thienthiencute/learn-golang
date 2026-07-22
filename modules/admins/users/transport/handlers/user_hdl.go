package handlers

import (
	"gocas/modules/admins/users/entity"
	"gocas/pkg/golangviet/templates"
	"gocas/views/admins/users"

	"github.com/gofiber/fiber/v2"
)

// trong day chi genre ra html, con api la xu ly

type UserUsc interface {
	DetailUserUsc(id int64) (*entity.User, error)
	ListRolesUsc() ([]entity.Role, error)
}

type userHdl struct {
	userUsc UserUsc
}

func NewUserHdl(userUsc UserUsc) *userHdl {
	return &userHdl{userUsc: userUsc}

}

func (h *userHdl) ListUserHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roles, _ := h.userUsc.ListRolesUsc()
		return templates.Render(c, users.Index(roles))
	}
}

func (h *userHdl) UpdateUserHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		data, err := h.userUsc.DetailUserUsc(int64(id))
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return templates.Render(c, users.Update(data))
	}
}
