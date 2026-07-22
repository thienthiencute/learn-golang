package responses

import "encoding/json"

type DetailUserGroupResp struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Status      int             `json:"status"`
	Permission  json.RawMessage `json:"permission"`
}
