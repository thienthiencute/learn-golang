package api

import (
	"fmt"
	"gocas/modules/admins/users_group/entity"
	"gocas/modules/admins/users_group/transport/requests"
	"gocas/modules/admins/users_group/transport/responses"

	"github.com/gofiber/fiber/v2"
)

type userGroupUsc interface {
	ListApiUserGroupUsc() (*[]responses.UserGroupResp, error)
	CreateApiUserGroupUsc(name, description string) error
	UpdateApiUserGroupUsc(id int64, name, description string, status int) error
	DetailApiUserGroupUsc(id int64) (*entity.UserGroup, error)
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

func (h *userGroupApi) CreateUserGroupApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.CreateUserGroupReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if req.Name == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
		}

		err := h.usc.CreateApiUserGroupUsc(req.Name, req.Description)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "User group created successfully",
		})
	}
}

func (h *userGroupApi) UpdateUserGroupApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.UpdateUserGroupReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if req.ID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID is required"})
		}
		if req.Name == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
		}

		err := h.usc.UpdateApiUserGroupUsc(req.ID, req.Name, req.Description, req.Status)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "User group updated successfully",
		})
	}
}

func (h *userGroupApi) DetailUserGroupApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		var id int64
		fmt.Sscanf(idStr, "%d", &id)

		data, err := h.usc.DetailApiUserGroupUsc(id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		if data == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User group not found"})
		}

		return c.JSON(fiber.Map{
			"data": data,
		})
	}
}
