package control

import (
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/restrict"
)

const (
	SuperAdminRole string = "SUPER_ADMIN"
	AdminRole      string = "ADMIN"
	UserRole       string = "USER"
)

var userRole *restrict.Role = &restrict.Role{
	ID:          UserRole,
	Description: "User is a standard user of the application.",
	Grants: map[string][]restrict.Action{
		models.MESSAGE_RESOURCE:      {restrict.Create, restrict.ReadAny, restrict.UpdateOwn, restrict.DeleteOwn},
		models.CONVERSATION_RESOURCE: {restrict.Create, restrict.ReadAny, restrict.DeleteOwn},
	},
}

var adminRole *restrict.Role = &restrict.Role{
	ID:          AdminRole,
	Description: "Admin can manage standard users.",
	Grants: map[string][]restrict.Action{
		models.USER_RESOURCE: {restrict.Create, restrict.ReadAny, restrict.UpdateAny, restrict.DeleteAny},
	},
	Parents: []string{UserRole},
}

var superAdminRole *restrict.Role = &restrict.Role{
	ID:          SuperAdminRole,
	Description: "SuperAdmin can manage all entities in the system.",
	Grants:      map[string][]restrict.Action{},
	Parents:     []string{AdminRole},
}

// Policy - describes Gochat's RBAC policy definition.
var Policy *restrict.PolicyDefinition = &restrict.PolicyDefinition{
	Resources: []string{
		models.USER_RESOURCE,
		models.CONVERSATION_RESOURCE,
		models.MESSAGE_RESOURCE,
	},
	Actions: []restrict.Action{
		restrict.Noop,
		restrict.Create,
		restrict.ReadAny,
		restrict.ReadOwn,
		restrict.UpdateAny,
		restrict.UpdateOwn,
		restrict.DeleteAny,
		restrict.DeleteOwn,
	},
	Roles: map[string]*restrict.Role{
		UserRole:       userRole,
		AdminRole:      adminRole,
		SuperAdminRole: superAdminRole,
	},
}
