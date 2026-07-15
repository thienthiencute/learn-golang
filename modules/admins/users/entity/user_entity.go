package entity

import (
	"encoding/json"
	"time"

	"github.com/teoit/gosctx/core"
)

type User struct {
	core.SQLModel
	Username     *string         `json:"username,omitempty"`
	Email        string          `json:"email"`
	FirstName    string          `json:"first_name"`
	LastName     string          `json:"last_name"`
	FullName     string          `json:"full_name"`
	Phone        *string         `json:"phone,omitempty"`
	Image        string          `json:"image"`
	Gender       int             `json:"gender"`
	Birthday     *time.Time      `json:"birthday,omitempty"`
	Status       int             `json:"status"`
	Permission   json.RawMessage `json:"permission"`
	IsPermission int             `json:"is_permission"`
	RoleID       int64           `json:"-" gorm:"column:role_id;"`
	RoleFake     *core.UID       `json:"role_id" gorm:"-"`
	CreatedBy    int64           `json:"created_by"`
	UpdatedBy    int64           `json:"updated_by"`
	// Role         *Role           `gorm:"foreignKey:RoleID"`
	UserCreated  *User           `gorm:"foreignKey:CreatedBy"`
	UserUpdated  *User           `gorm:"foreignKey:UpdatedBy"`
}

func (User) TableName() string {
	return "admins.users"
}

// func (u *User) Mask() {
// 	u.SQLModel.Mask(gocommon.MASK_TYPE_USER)
// 	uid := core.NewUID(uint32(u.RoleID), gocommon.MASK_TYPE_ROLE, 1)
// 	u.RoleFake = &uid
// }