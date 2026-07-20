package api

import (
	"context"
	"gocas/modules/admins/users_group/entity"
	"gocas/modules/admins/users_group/transport/requests"
	"gocas/modules/admins/users_group/transport/responses"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx/core"
)

type userGroupApiUsc interface {
	ListApiUserGroupUsc() (*[]responses.UserGroupResp, error)
	CreateApiUserGroupUsc(ctx context.Context, req *requests.UserGroupCreation) error
	UpdateApiUserGroupUsc(ctx context.Context, req *requests.UserGroupUpdateReq) error
	DetailApiUserGroupUsc(ctx context.Context, req *requests.UserGroupDetailReq) (*entity.UserGroup, error)
	DeleteApiUserGroupUsc(ctx context.Context, id int64) error
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

		err := h.userGroupApiUsc.CreateApiUserGroupUsc(c.UserContext(), &req)
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
		var req requests.UserGroupUpdateReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		err := h.userGroupApiUsc.UpdateApiUserGroupUsc(c.UserContext(), &req)
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
		var req requests.UserGroupDetailReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		data, err := h.userGroupApiUsc.DetailApiUserGroupUsc(c.UserContext(), &req)
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
		var req requests.UserGroupDeleteReq

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"details": fiber.Map{"msg": "Invalid request body"}})
		}

		if validationErr := req.Validation(); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		for _, idStr := range req.Ids {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"details": fiber.Map{"msg": "Invalid ID format"}})
			}
			err = h.userGroupApiUsc.DeleteApiUserGroupUsc(c.UserContext(), id)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"details": fiber.Map{"msg": err.Error()}})
			}
		}

		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"msg": "User group(s) deleted successfully",
			},
		})
	}
}
