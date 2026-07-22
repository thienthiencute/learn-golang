package handlers

import (
	"gocas/modules/admins/users_group/entity"
	"gocas/pkg/golangviet/templates"
	"gocas/views/admins/users_group"
	errorpages "gocas/views/pages"

	"github.com/gofiber/fiber/v2"
)

// userGroupUsc: interface mà handler phụ thuộc vào (chỉ cần method này để render trang Update)
// -> handler không quan tâm usecase implement thế nào, chỉ cần đúng contract
type userGroupUsc interface {
	FindOneUserGroupUsc(id int) (*entity.UserGroup, error)
}

// userGroupHdl: struct chứa dependency (usecase) để xử lý request
type userGroupHdl struct {
	userGroupUsc userGroupUsc
}

// NewUserGroupHdl: hàm khởi tạo (constructor), nhận usecase từ ngoài vào (Dependency Injection)
func NewUserGroupHdl(userGroupUsc userGroupUsc) *userGroupHdl {
	return &userGroupHdl{userGroupUsc: userGroupUsc}
}

// ListUserGroupHdl: render trang danh sách nhóm quyền
// -> không cần lấy data ở đây vì có thể trang này tự gọi API (list) qua JS/AJAX để load data
func (h *userGroupHdl) ListUserGroupHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return templates.Render(c, users_group.Index()) // chỉ render khung view, không truyền data
	}
}

// UpdateUserGroupHdl: render trang sửa nhóm quyền theo id, có lấy sẵn data để hiển thị form
func (h *userGroupHdl) UpdateUserGroupHdl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Lấy id từ URL param (/update/:id) và ép kiểu int
		id, err := c.ParamsInt("id")

		// Validate: nếu id không parse được hoặc <= 0 -> trả trang 404
		if err != nil || id <= 0 {
			return templates.Render(c, errorpages.NotFound404("/admins/roles/list"))
		}

		// Gọi usecase lấy thông tin chi tiết nhóm quyền theo id
		detailUserGroup, err := h.userGroupUsc.FindOneUserGroupUsc(id)

		// Nếu không tìm thấy (hoặc lỗi DB) -> trả trang 404
		if err != nil {
			return templates.Render(c, errorpages.NotFound404("/admins/roles/list"))
		}

		// Render trang Update, truyền sẵn data vào để hiển thị (fill form)
		return templates.Render(c, users_group.Update(detailUserGroup))
	}
}