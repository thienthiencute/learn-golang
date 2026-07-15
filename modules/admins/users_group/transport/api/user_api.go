package api

import (
	"fmt"
	"gocas/modules/admins/users_group/transport/responses"

	"github.com/gofiber/fiber/v2"
)

type userGroupUsc interface {
	ListApiUserGroupUsc() (*[]responses.UserGroupResp, error)
}

type userGroupApi struct {
	usc userGroupUsc
}

func NewUserGroupApi(usc userGroupUsc) *userGroupApi {
	return &userGroupApi{usc: usc}
}

func (h *userGroupApi) ListUserGroupApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("123")
		// Gọi UseCase để lấy dữ liệu thực tế từ Database
		data, err := h.usc.ListApiUserGroupUsc()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Trả về JSON với key "data" cho DataTables
		return c.JSON(fiber.Map{
			"data": data,
		})


		
	}
}
