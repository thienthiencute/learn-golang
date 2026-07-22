package mapping

import "gocas/modules/admins/users_group/transport/requests"

func MapperCondListDatableUserGroup(req *requests.ListDatatableUserGroupReq) map[string]interface{} {
	data := map[string]interface {
	}{}
	if req.SearchName != "" {
		data["name ILIKE ?"] = "%" + req.SearchName + "%"
	}
	return data
}
