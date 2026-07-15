package entity

import (
	"encoding/json"

	"github.com/teoit/gosctx/core"
)

type UserGroup struct {
	core.SQLModel
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Status       int             `json:"status"`
	Permission   json.RawMessage `json:"permission"`
	CreatedBy    int64           `json:"created_by"`
	UpdatedBy    int64           `json:"updated_by"`
}

func (UserGroup) TableName() string {
	return "admins.users_group"
}
