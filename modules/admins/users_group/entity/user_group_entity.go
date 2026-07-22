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

}

func (UserGroup) TableName() string {
	return "admins.users_group"
}
