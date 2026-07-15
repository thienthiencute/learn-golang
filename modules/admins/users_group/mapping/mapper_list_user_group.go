package mapping

import (
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

		status := "Hien"
		if userGroup.Status == 2 {
			status = "An"
		}

		str := strconv.FormatInt(userGroup.ID, 10)
		result = append(result, responses.UserGroupResp{
			ID:          userGroup.ID,
			Name:        userGroup.Name,
			Description: userGroup.Description,
			Status:      status,
			Custom: `
                        <button class="btn btn-sm btn-info edit-btn" data-id="` + str + `">Edit</button>
                        <button class="btn btn-sm btn-danger delete-btn" data-id="` + str + `">Delete</button>
                    `,
		})

	}

	//

	return &result
}
