package utils

type Filters struct {
	Preloads  *map[string]interface{}
	Columns   *[]string
	Conds     *map[string]interface{}
	UpdateMap *map[string]interface{}
	PageSize  int
	Offset    int
	OrderBy   *[]string
}

type UpdateFilters struct {
	Columns    *[]string
	Conds      *map[string]interface{}
	FieldIN    *map[string]interface{}
	FieldNotIN *map[string]interface{}
	FieldLike  *map[string]string
	UpdateMap  *map[string]interface{}
}
