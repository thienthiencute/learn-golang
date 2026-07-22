package mapping

import (
	"fmt"
	"gocas/modules/admins/common/consts"
	"gocas/modules/admins/users/entity"
	"gocas/modules/admins/users/transport/responses"
	"strconv"
)

func MapperListUser(data *[]entity.User) *[]responses.UserResp {
	var result []responses.UserResp
	for _, user := range *data {
		statusStr, ok := consts.MapStatusIntToString[user.Status]
		if !ok {
			statusStr = consts.STATUS_ACTIVE_STR
		}

		strID := strconv.FormatInt(user.ID, 10)
		
		avatar := user.Image
		if avatar == "" {
			avatar = "/static/images/users/user-dummy-img.jpg" 
		}
		fullName := user.LastName + " " + user.FirstName
		fullNameHTML := fmt.Sprintf(`
			<div class="d-flex align-items-center">
				<img src="%s" class="me-3 rounded-circle avatar-xs" alt="user-pic">
				<div class="flex-grow-1">
					<h6 class="m-0">%s</h6>
					<span class="fs-11 mb-0 text-muted">%s</span>
				</div>
			</div>
		`, avatar, fullName, user.Email)

		auditInfoHTML := ""
		if user.UserCreated != nil {
			auditInfoHTML = fmt.Sprintf(`<b>Tạo:</b> %s - %s`, user.UserCreated.FullName, user.CreatedAt.Format("2006-01-02 15:04"))
		} else {
			auditInfoHTML = fmt.Sprintf(`<b>Tạo:</b> System - %s`, user.CreatedAt.Format("2006-01-02 15:04"))
		}

		roleName := "ADMIN"
		if user.Role != nil {
			roleName = user.Role.Name
		}
		roleHTML := fmt.Sprintf(`<span class="text-primary">%s</span>`, roleName)

		result = append(result, responses.UserResp{
			ID:            user.ID,
			FullNameHTML:  fullNameHTML,
			RoleHTML:      roleHTML,
			Status:        statusStr,
			AuditInfoHTML: auditInfoHTML,
			Custom: `
                    <ul class="list-inline hstack gap-2 mb-0 d-flex">
                        <li class="list-inline-item" data-bs-toggle="tooltip" data-bs-trigger="hover" data-bs-placement="top" title="Xem chi tiết">
                            <a href="/admins/users/view/` + strID + `" class="text-primary d-inline-block">
                                <i class="ri-eye-fill fs-16"></i>
                            </a>
                        </li>
                        <li class="list-inline-item item_edit" data-bs-toggle="tooltip" data-bs-trigger="hover" data-bs-placement="top" title="Chỉnh sửa">
                            <a class="edit-item-btn" href="/admins/users/update/` + strID + `">
                                <i class="ri-pencil-fill align-bottom"></i>
                            </a>
                        </li>    
                        <li class="list-inline-item item_delete" data-bs-toggle="tooltip" data-bs-trigger="hover" data-bs-placement="top" title="Xóa">
                            <a id="delete-item-btn" href="#delete_modal" class="text-danger" data-bs-toggle="modal">
                                <i class="ri-delete-bin-5-fill align-bottom"></i>
                            </a>
                        </li>   
                    </ul>
                    `,
		})
	}

	if result == nil {
		result = []responses.UserResp{}
	}

	return &result
}
