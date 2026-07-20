package mapping

import (
	"gocas/modules/admins/common/consts"
	"gocas/modules/admins/users_group/entity"
	"gocas/modules/admins/users_group/transport/responses"
	"strconv"
)

func MapperListUserGroup(data *[]entity.UserGroup) *[]responses.UserGroupResp {
	var result []responses.UserGroupResp
	for _, userGroup := range *data {
		// status := ""
		// if userGroup.Status == 1 {
		// 	status = "Hien"
		// } else{
		// 	status = "An"
		// }

		statusStr, ok := consts.MapStatusIntToString[userGroup.Status]
		if !ok {
			statusStr = consts.STATUS_ACTIVE_STR
		}

		str := strconv.FormatInt(userGroup.ID, 10)
		result = append(result, responses.UserGroupResp{
			ID:          userGroup.ID,
			Name:        userGroup.Name,
			Description: userGroup.Description,
			Status:      statusStr,
			Custom: `
                    <ul class="list-inline hstack gap-2 mb-0 d-flex">

                        <li class="list-inline-item item_edit" data-bs-toggle="tooltip" data-bs-trigger="hover" data-bs-placement="top" aria-label="Edit" data-bs-original-title="Chỉnh sửa" data-key="t-edit">
                            <a class="edit-item-btn" href="/admins/roles/update/` + str + `">
                                <i class="ri-pencil-fill align-bottom"></i>
                            </a>
                        </li>    
                        <li class="list-inline-item item_delete" data-bs-toggle="tooltip" data-bs-trigger="hover" data-bs-placement="top" aria-label="Delete" data-bs-original-title="Xóa" data-key="t-delete">
                            <a id="delete-item-btn" href="#delete_modal" class="text-danger" data-bs-toggle="modal">
                                <i class="ri-delete-bin-5-fill align-bottom"></i>
                            </a>
                        </li>   
                    </ul>
                    `,
		})

	}

	//

	return &result
}
