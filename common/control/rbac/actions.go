package rbac

// ActionType - describes an action type a role's member
// can take against a resource.
// type ActionType string

type Action string

const (
	Noop      Action = ""           // Noop - empty action.
	Create    Action = "CREATE"     // Create - action for creating a resource.
	ReadAny   Action = "READ_ANY"   // ReadAny - action for reading any resource of given type.
	ReadOwn   Action = "READ_OWN"   // ReadOwn - action for reading own resource of given type.
	UpdateAny Action = "UPDATE_ANY" // UpdateAny - action for updating any resource of given type.
	UpdateOwn Action = "UPDATE_OWN" // UpdateOwn - action for updating own resource of given type.
	DeleteAny Action = "DELETE_ANY" // DeleteAny - action for deleting any resource of given type.
	DeleteOwn Action = "DELETE_OWN" // DeleteOwn - action for deleting own resource of given type.
)

// List of basic actions.
// const (
// 	Noop      ActionType = ""           // Noop - empty action.
// 	Create    ActionType = "CREATE"     // Create - action for creating a resource.
// 	ReadAny   ActionType = "READ_ANY"   // ReadAny - action for reading any resource of given type.
// 	ReadOwn   ActionType = "READ_OWN"   // ReadOwn - action for reading own resource of given type.
// 	UpdateAny ActionType = "UPDATE_ANY" // UpdateAny - action for updating any resource of given type.
// 	UpdateOwn ActionType = "UPDATE_OWN" // UpdateOwn - action for updating own resource of given type.
// 	DeleteAny ActionType = "DELETE_ANY" // DeleteAny - action for deleting any resource of given type.
// 	DeleteOwn ActionType = "DELETE_OWN" // DeleteOwn - action for deleting own resource of given type.
// )

// ActionModifiler - describes a target of an action
// type ActionModifier string

// const (
// 	Any    ActionModifier = "ANY"
// 	Own    ActionModifier = "OWN"
// 	Self   ActionModifier = "SELF"
// 	PartOf ActionModifier = "PART_OF"
// )

// Action - describes an action a role's member can
// take against a resource in regard of given modifier.
// type Action struct {
// 	Type     ActionType
// 	Modifier ActionModifier
// }
