package api

import (
	"gocas/modules/admins/users_group/entity"
	"gocas/modules/admins/users_group/transport/requests"
	"gocas/modules/admins/users_group/transport/responses"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx/core"
)

type userGroupApiUsc interface {
	ListApiUserGroupUsc() (*[]responses.UserGroupResp, error)
	CreateApiUserGroupUsc(name, description string, status int) error
	UpdateApiUserGroupUsc(id int64, name, description string, status int) error
	DetailApiUserGroupUsc(id int64) (*entity.UserGroup, error)
	DeleteApiUserGroupUsc(id int64) error
}

type userGroupApi struct {
	userGroupApiUsc userGroupApiUsc
}

func NewUserGroupApi(userGroupApiUsc userGroupApiUsc) *userGroupApi {
	return &userGroupApi{userGroupApiUsc: userGroupApiUsc}
}

func (h *userGroupApi) ListUserGroupApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := h.userGroupApiUsc.ListApiUserGroupUsc()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"data": data,
		})
	}
}

func (h *userGroupApi) CreateUserGroupApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.UserGroupCreation
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		err := h.userGroupApiUsc.CreateApiUserGroupUsc(req.Name, req.Description, req.Status)
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
		var req requests.UserGroupUpdate
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		err := h.userGroupApiUsc.UpdateApiUserGroupUsc(int64(req.ID), req.Name, req.Description, req.Status)
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
		var req requests.UserGroupDetail
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		data, err := h.userGroupApiUsc.DetailApiUserGroupUsc(req.ID)
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

func (h *userGroupApi) DeleteUserGroupApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.UserGroupDetail
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		err := h.userGroupApiUsc.DeleteApiUserGroupUsc(req.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "User group deleted successfully",
		})
	}
}
