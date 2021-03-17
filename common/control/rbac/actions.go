package rbac

// Action - describes an action a role's member
// can take against a resource.
type Action string

const (
	// Noop - empty action.
	Noop Action = ""
	// Create - action for creating a resource.
	Create Action = "CREATE"
	// Read - action for reading a resource.
	Read Action = "READ"
	// Update - action for updating a resource.
	Update Action = "UPDATE"
	// Delete - action for deleting a resource.
	Delete Action = "DELETE"
	// CRUD - action for all four CRUD operations.
	CRUD Action = "CRUD"
)
