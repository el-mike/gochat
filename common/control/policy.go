package control

import (
	"github.com/el-Mike/gochat/common/control/rbac"
	"github.com/el-Mike/gochat/models"
)

const (
	SuperAdminRole string = "SUPER_ADMIN"
	AdminRole      string = "ADMIN"
	UserRole       string = "USER"
)

var userRole *rbac.Role = &rbac.Role{
	ID:          UserRole,
	Description: "User is a standard user of the application.",
	Grants: map[string][]rbac.Action{
		models.MESSAGE_RESOURCE:      {rbac.Create, rbac.ReadAny, rbac.UpdateOwn, rbac.DeleteOwn},
		models.CONVERSATION_RESOURCE: {rbac.Create, rbac.ReadAny, rbac.DeleteOwn},
	},
}

var adminRole *rbac.Role = &rbac.Role{
	ID:          AdminRole,
	Description: "Admin can manage standard users.",
	Grants: map[string][]rbac.Action{
		models.USER_RESOURCE: {rbac.Create, rbac.ReadAny, rbac.UpdateAny, rbac.DeleteAny},
	},
	Parents: []string{UserRole},
}

var superAdminRole *rbac.Role = &rbac.Role{
	ID:          SuperAdminRole,
	Description: "SuperAdmin can manage all entities in the system.",
	Grants:      map[string][]rbac.Action{},
	Parents:     []string{AdminRole},
}

// Policy - describes Gochat's RBAC policy definition.
var Policy *rbac.PolicyDefinition = &rbac.PolicyDefinition{
	Resources: []string{
		models.USER_RESOURCE,
		models.CONVERSATION_RESOURCE,
		models.MESSAGE_RESOURCE,
	},
	Actions: []rbac.Action{
		rbac.Noop,
		rbac.Create,
		rbac.ReadAny,
		rbac.ReadOwn,
		rbac.UpdateAny,
		rbac.UpdateOwn,
		rbac.DeleteAny,
		rbac.DeleteOwn,
	},
	Roles: map[string]*rbac.Role{
		UserRole:       userRole,
		AdminRole:      adminRole,
		SuperAdminRole: superAdminRole,
	},
}
