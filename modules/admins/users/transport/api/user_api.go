package api

import (
	"gocas/modules/admins/users/usecase"
	"gocas/modules/admins/users/transport/requests"
	"gocas/modules/admins/users/entity"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx/core"
)

type userApi struct {
	userApiUsc usecase.UserApiUsc
}

func NewUserApi(userApiUsc usecase.UserApiUsc) *userApi {
	return &userApi{userApiUsc: userApiUsc}
}

func (h *userApi) ListUserApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := h.userApiUsc.ListApiUserUsc(c.UserContext())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"data": data,
		})
	}
}

func (h *userApi) CreateUserApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.CreateUserReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		
		user := &entity.User{
			Email:     req.Email,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			RoleID:    req.RoleID,
			Status:    req.Status,
		}

		_, err := h.userApiUsc.CreateApiUserUsc(c.UserContext(), user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

func (h *userApi) UpdateStatusUserApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.UserUpdateStatusReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		err := h.userApiUsc.UpdateStatusApiUserUsc(c.UserContext(), &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "User status updated successfully",
		})
	}
}

func (h *userApi) UpdateUserApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.UpdateUserReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		err := h.userApiUsc.UpdateApiUserUsc(c.UserContext(), &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "User updated successfully",
		})
	}
}

func (h *userApi) DeleteUserApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.DeleteUserReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		err := h.userApiUsc.DeleteApiUserUsc(c.UserContext(), &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "User deleted successfully",
		})
	}
}
