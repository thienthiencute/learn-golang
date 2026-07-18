package consts

const (
	STATUS_ACTIVE   = 1
	STATUS_INACTIVE = 2

)

var (
	StatusUserGroupToViewCreate = []int{
		STATUS_ACTIVE,
		STATUS_INACTIVE}
	MapStatusUserGroup = map[int]string{
		STATUS_ACTIVE:   "Hien",
		STATUS_INACTIVE: "An",
	}
)