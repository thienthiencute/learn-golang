package responses

type UserResp struct {
	ID            int64  `json:"id"`
	FullNameHTML  string `json:"full_name_html"`
	RoleHTML      string `json:"role_html"`
	Status        string `json:"status"`
	AuditInfoHTML string `json:"audit_info_html"`
	Custom        string `json:"custom"`
}
