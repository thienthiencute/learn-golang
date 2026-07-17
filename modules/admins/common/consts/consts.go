package consts

const (
	STATUS_ACTIVE   = 1
	STATUS_INACTIVE = 2

)

var (
	StatusUserGroupToViewCreate = []int{STATUS_ACTIVE, STATUS_INACTIVE}
	MapStatusStr = map[int]string{
		STATUS_ACTIVE:   "ACTIVE",
		STATUS_INACTIVE: "INACTIVE",
	}
	MapStatusInt = map[string]int{
		"ACTIVE":   STATUS_ACTIVE,
		"INACTIVE": STATUS_INACTIVE,
	}
)