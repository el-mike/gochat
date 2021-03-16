package control

import (
	"github.com/google/uuid"
)

// ContextUserKey - defines the key currentUser will be saved under in current context.
var ContextUserKey = "currentUser"

// ContextUser - struct defining user assigned to current context.
type ContextUser struct {
	ID       uuid.UUID `json:"id"`
	AuthUUID uuid.UUID `json:"authUUID"`
	Email    string    `json:"email"`
}
