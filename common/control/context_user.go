package control

import (
	"github.com/el-Mike/gochat/common/control/rbac"
	"github.com/google/uuid"
)

// ContextUserKey - defines the key currentUser will be saved under in current context.
var ContextUserKey = "currentUser"

// ContextUser - struct defining user assigned to current context.
type ContextUser struct {
	ID       uuid.UUID    `json:"id"`
	AuthUUID uuid.UUID    `json:"authUUID"`
	Email    string       `json:"email"`
	Roles    *[]rbac.Role `json:"roles"`
}

func (cu *ContextUser) Role() *rbac.Role {
	return &rbac.Role{
		ID:          "TEST_ROLE",
		Description: "TEST_DESC",
		Grants:      rbac.GrantsMap{},
		Parents:     []string{},
	}
}
