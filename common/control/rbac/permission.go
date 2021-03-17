package rbac

// Permission - definition of an action or set of actions
// granted to a given Role.
type Permission struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Actions     []Action `json:"actions"`
}
