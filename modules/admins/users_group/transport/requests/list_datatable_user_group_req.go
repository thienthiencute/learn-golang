package requests

type ListDatatableUserGroupReq struct {
	Draw    int      `form:"draw"`
	Start   int      `form:"start"`
	Length  int      `form:"length"`
	Search  Search   `form:"search"`
	Order   []Order  `form:"order"`
	Columns []Column `form:"columns"`
	SearchName string `form:"search_name"`

}

type Search struct {
	Value string `form:"search[value]"`
}

type Order struct {
	Column int    `form:"column"`
	Dir    string `form:"dir"`
}

type Column struct {
	Data string `form:"data"`
}
