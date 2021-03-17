package rbac

// AccessManager - manages all of the defined Permissions and Roles,
// provides an interface to perform authorization checks and add/remove
// Permissions and Roles.
type AccessManager struct {
	Permissions []*Permission `json:"permissions"`
	Roles       []*Role       `json:"roles"`
}

// NewAccessManager - returns a new AccessManager instance.
func NewAccessManager() *AccessManager {
	return &AccessManager{}
}

func (am *AccessManager) IsGranted(role *Role, permissions ...*Permission) bool {
	return true
}
