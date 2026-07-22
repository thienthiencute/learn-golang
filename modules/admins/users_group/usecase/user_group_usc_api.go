package usecase

import (
	"context"
	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users_group/entity"
	"gocas/modules/admins/users_group/mapping"
	"gocas/modules/admins/users_group/transport/requests"
	"gocas/modules/admins/users_group/transport/responses"
	"gocas/pkg/golangviet/utils"
)

// userGroupApi: interface local (consumer-defined) mà Usecase cần từ tầng Repository.
// Usecase chỉ phụ thuộc interface này -> không biết chi tiết DB bên dưới (GORM, raw SQL...),
// dễ mock khi viết unit test.
type userGroupApi interface {
	ListUserGroupRepo() (*[]entity.UserGroup, error)
	FindUserGroupRepo(ctx context.Context, filter *utils.Filters) ([]*entity.UserGroup, error)
	CreateUserGroupRepo(data *entity.UserGroup) error
	UpdateUserGroupRepo(id int64, data map[string]interface{}) error
	GetUserGroupByIdRepo(id int64) (*entity.UserGroup, error)
	DeleteUserGroupRepo(ids []int64) error
}

// userGroupApiUsc: struct chứa dependency Repo, implement business logic cho module user group (phía API)
type userGroupApiUsc struct {
	userGroupApi userGroupApi
}

// NewUserGroupApiUsc: constructor - nhận Repo từ ngoài vào (Dependency Injection)
func NewUserGroupApiUsc(userGroupApi userGroupApi) *userGroupApiUsc {
	return &userGroupApiUsc{userGroupApi: userGroupApi}
}

// ListApiUserGroupUsc: lấy danh sách nhóm quyền, map từ entity (DB model) sang response DTO
func (u *userGroupApiUsc) ListApiUserGroupUsc() (*[]responses.UserGroupResp, error) {
	_, err := u.userGroupApi.ListUserGroupRepo()
	if err != nil {
		// Không tìm thấy record -> coi là "danh sách rỗng", không phải lỗi thật sự
		// (trả nil, nil thay vì trả lỗi ra ngoài, để tầng trên xử lý như list rỗng)
		if err == errs.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	// mapping: convert []entity.UserGroup (DB model) -> []responses.UserGroupResp (response DTO)
	// tách biệt cấu trúc DB khỏi cấu trúc trả ra client
	// result := mapping.MapperListUserGroup(data)

	return nil, nil
}

func (u *userGroupApiUsc) ListDataTableApiUserGroupUsc(ctx context.Context, req *requests.ListDatatableUserGroupReq) (*responses.ListDatatableUserGroupResp, error) {
	limit := req.Length
	conds := mapping.MapperCondListDatableUserGroup(req)
	// orderColumn := req.Order[0].Column
	// orderDir := req.Order[0].Dir
	// order, ok := common.ORDER_DATATABLE[orderColumn]
	// if !ok {
	// 	order = cenum.ORDER_DEFAULT
	// } else {
	// 	order = order + " " + orderDir
	// }
	filter := utils.Filters{
		Conds:    &conds,
		Offset:   req.Start,
		PageSize: limit,
		// OrderBy:  &[]string{order},
	}
	data, err := u.userGroupApi.FindUserGroupRepo(ctx, &filter)
	if err != nil {
		if err == errs.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	result := mapping.MapperListUserGroup(data)

	userDatatable := &responses.ListDatatableUserGroupResp{
		Draw:            req.Draw,
		RecordsTotal:    20,
		RecordsFiltered: 10,
		Data:            result,
	}
	return userDatatable, nil
}

// CreateApiUserGroupUsc: tạo mới 1 nhóm quyền
// Convert request DTO (req) -> entity (DB model) trước khi gọi Repo lưu xuống DB
func (u *userGroupApiUsc) CreateApiUserGroupUsc(ctx context.Context, req *requests.UserGroupCreation) error {
	userGroup := &entity.UserGroup{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
	}

	err := u.userGroupApi.CreateUserGroupRepo(userGroup)
	if err != nil {
		return err
	}

	return nil
}

// UpdateApiUserGroupUsc: cập nhật toàn bộ 3 field (name, description, status) của 1 nhóm quyền
// Dùng map[string]interface{} thay vì entity -> phù hợp khi Repo dùng GORM.Updates(map) để chỉ update
// đúng các field chỉ định (tránh ghi đè field khác không mong muốn, nhưng cũng làm mất type-safety)
func (u *userGroupApiUsc) UpdateApiUserGroupUsc(ctx context.Context, req *requests.UserGroupUpdateReq) error {
	data := map[string]interface{}{
		"name":        req.Name,
		"description": req.Description,
		"status":      req.Status,
	}

	err := u.userGroupApi.UpdateUserGroupRepo(int64(req.ID), data)
	if err != nil {
		return err
	}

	return nil
}

// UpdateStatusApiUserGroupUsc: chỉ update riêng field "status"
// Tái sử dụng chung hàm Repo UpdateUserGroupRepo (nhận map) với Update ở trên,
// chỉ khác là map chỉ chứa 1 field "status" -> GORM chỉ generate UPDATE cho field này
func (u *userGroupApiUsc) UpdateStatusApiUserGroupUsc(ctx context.Context, req *requests.UserGroupUpdateStatusReq) error {
	data := map[string]interface{}{
		"status": req.Status,
	}

	err := u.userGroupApi.UpdateUserGroupRepo(int64(req.ID), data)
	if err != nil {
		return err
	}

	return nil
}

// DetailApiUserGroupUsc: lấy chi tiết 1 nhóm quyền theo ID
// Trả về DTO (responses.DetailUserGroupResp) để ẩn các field nội bộ của DB (created_at, updated_at...)
func (u *userGroupApiUsc) DetailApiUserGroupUsc(ctx context.Context, req *requests.UserGroupDetailReq) (*responses.DetailUserGroupResp, error) {
	data, err := u.userGroupApi.GetUserGroupByIdRepo(req.ID)
	if err != nil {
		if err == errs.ErrRecordNotFound {
			return nil, nil // không tìm thấy -> trả nil, nil để Api tầng trên tự check và trả 404
		}
		return nil, err
	}

	result := mapping.MapperDetailUserGroup(data)
	return result, nil
}

// DeleteApiUserGroupUsc: xóa nhiều nhóm quyền theo danh sách id (bulk delete)
// Gọi Repo 1 lần duy nhất, xóa hết trong 1 query -> atomic hơn, nhanh hơn loop từng id.
func (u *userGroupApiUsc) DeleteApiUserGroupUsc(ctx context.Context, ids []int64) error {
	err := u.userGroupApi.DeleteUserGroupRepo(ids)
	if err != nil {
		return err
	}
	return nil
}
