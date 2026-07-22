package responses

type ListDatatableUserGroupResp struct {
	Draw            int              `json:"draw"`
	RecordsTotal    int64            `json:"recordsTotal"`
	RecordsFiltered int64            `json:"recordsFiltered"`
	Data            *[]UserGroupResp `json:"data"`
}
