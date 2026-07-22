package api

import (
	"context"
	"fmt"
	"gocas/modules/admins/users_group/transport/requests"
	"gocas/modules/admins/users_group/transport/responses"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/teoit/gosctx/core"
)

// userGroupApiUsc: interface local (consumer-defined) mà Api handler cần từ tầng Usecase.
// Api chỉ phụ thuộc vào interface này -> dễ mock test, không lệ thuộc implementation cụ thể.
type userGroupApiUsc interface {
	ListApiUserGroupUsc() (*[]responses.UserGroupResp, error)
	ListDataTableApiUserGroupUsc(ctx context.Context, req *requests.ListDatatableUserGroupReq) (*responses.ListDatatableUserGroupResp, error)
	CreateApiUserGroupUsc(ctx context.Context, req *requests.UserGroupCreation) error
	UpdateApiUserGroupUsc(ctx context.Context, req *requests.UserGroupUpdateReq) error
	UpdateStatusApiUserGroupUsc(ctx context.Context, req *requests.UserGroupUpdateStatusReq) error
	DetailApiUserGroupUsc(ctx context.Context, req *requests.UserGroupDetailReq) (*responses.DetailUserGroupResp, error)
	DeleteApiUserGroupUsc(ctx context.Context, ids []int64) error
}

// userGroupApi: struct chứa dependency Usecase, implement các fiber.Handler cho route /api/admins/roles
type userGroupApi struct {
	userGroupApiUsc userGroupApiUsc
}

// NewUserGroupApi: constructor - nhận Usecase từ ngoài vào (Dependency Injection)
func NewUserGroupApi(userGroupApiUsc userGroupApiUsc) *userGroupApi {
	return &userGroupApi{userGroupApiUsc: userGroupApiUsc}
}

// ListUserGroupApi godoc
// @Summary      Danh sách nhóm quyền
// @Description  Lấy danh sách tất cả các nhóm quyền
// @Tags         roles
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /api/admins/roles/list [post]
// func (h *userGroupApi) ListUserGroupApi() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		data, err := h.userGroupApiUsc.ListApiUserGroupUsc()
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 		}

// 		return c.JSON(fiber.Map{
// 			"data": data,
// 		})
// 	}
// }

func (h *userGroupApi) ListUserGroupApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.ListDatatableUserGroupReq
		fmt.Println("==================================")
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		fmt.Printf("%+v\n", req)

		data, _ := h.userGroupApiUsc.ListDataTableApiUserGroupUsc(c.Context(), &req)

		return c.JSON(data)
	}
}

// CreateUserGroupApi godoc
// @Summary      Tạo mới nhóm quyền
// @Description  Tạo mới một nhóm quyền (Role)
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        request body requests.UserGroupCreation true "Thông tin nhóm quyền"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/admins/roles/create [post]
func (h *userGroupApi) CreateUserGroupApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.UserGroupCreation
		if err := c.BodyParser(&req); err != nil { // parse JSON body vào struct req
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil { // validate field bắt buộc, format...
			return core.ReturnErrsForApi(c, validationErr) // helper format lỗi validation trả về client
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

// UpdateUserGroupApi godoc
// @Summary      Cập nhật nhóm quyền
// @Description  Cập nhật toàn bộ thông tin 1 nhóm quyền (name, description, status)
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        request body requests.UserGroupUpdateReq true "Thông tin cập nhật nhóm quyền"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/admins/roles/update [post]
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

// UpdateStatusUserGroupApi godoc
// @Summary      Cập nhật trạng thái nhóm quyền
// @Description  Chỉ cập nhật riêng field status (kích hoạt/vô hiệu hóa nhóm quyền)
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        request body requests.UserGroupUpdateStatusReq true "Trạng thái mới"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/admins/roles/update-status [patch]
func (h *userGroupApi) UpdateStatusUserGroupApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.UserGroupUpdateStatusReq
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		err := h.userGroupApiUsc.UpdateStatusApiUserGroupUsc(c.UserContext(), &req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{
			"message": "User group status updated successfully",
		})
	}
}

// DetailUserGroupApi godoc
// @Summary      Chi tiết nhóm quyền
// @Description  Lấy thông tin chi tiết 1 nhóm quyền theo request body (không dùng URL param :id)
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        request body requests.UserGroupDetailReq true "ID của nhóm quyền"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/admins/roles/detail [post]
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
		if data == nil { // usecase không lỗi nhưng cũng không có data -> nghĩa là không tìm thấy
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User group not found"})
		}

		return c.JSON(fiber.Map{
			"data": data,
		})
	}
}

// DeleteUserGroupApi godoc
// @Summary      Xóa nhóm quyền
// @Description  Xóa nhiều nhóm quyền theo mảng ID
// @Tags         roles
// @Accept       json
// @Produce      json
// @Param        request body requests.UserGroupDeleteReq true "Danh sách ID cần xóa"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/admins/roles/delete [delete]
func (h *userGroupApi) DeleteUserGroupApi() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req requests.UserGroupDeleteReq

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"details": fiber.Map{"msg": "Invalid request body"}})
		}

		if validationErr := req.Validation(c.Context()); validationErr != nil {
			return core.ReturnErrsForApi(c, validationErr)
		}

		ids := make([]int64, 0, len(req.Ids))
		for _, idStr := range req.Ids {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"details": fiber.Map{"msg": "Invalid ID format: " + idStr}})
			}
			ids = append(ids, id)
		}

		if err := h.userGroupApiUsc.DeleteApiUserGroupUsc(c.UserContext(), ids); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"details": fiber.Map{"msg": err.Error()}})
		}

		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"msg": "User group(s) deleted successfully",
			},
		})
	}
}
