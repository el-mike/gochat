package control

import (
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/restrict"
)

const (
	SuperAdminRole = "SUPER_ADMIN"
	AdminRole      = "ADMIN"
	UserRole       = "USER"
)

const (
	CreateAction = "create"

	ReadAction = "read"

	UpdateAction    = "update"
	UpdateOwnAction = "updateOwn"

	DeleteAction    = "delete"
	DeleteOwnAction = "deleteOwn"
)

const (
	AccessOwnPreset = "accessOwn"
)

var userRole = &restrict.Role{
	ID:          UserRole,
	Description: "User is a standard user of the application.",
	Grants: restrict.GrantsMap{
		models.MESSAGE_RESOURCE: {
			&restrict.Permission{Action: CreateAction},
			&restrict.Permission{Action: ReadAction, Preset: AccessOwnPreset},
			&restrict.Permission{Action: UpdateOwnAction, Preset: AccessOwnPreset},
			&restrict.Permission{Action: DeleteOwnAction, Preset: AccessOwnPreset},
		},
		models.CONVERSATION_RESOURCE: {
			&restrict.Permission{Action: CreateAction},
			&restrict.Permission{Action: ReadAction},
			&restrict.Permission{Action: DeleteAction, Preset: AccessOwnPreset},
		},
	},
}

var adminRole *restrict.Role = &restrict.Role{
	ID:          AdminRole,
	Description: "Admin can manage standard users.",
	Grants: restrict.GrantsMap{
		models.USER_RESOURCE: {
			&restrict.Permission{Action: CreateAction},
			&restrict.Permission{Action: ReadAction, Preset: AccessOwnPreset},
			&restrict.Permission{Action: UpdateAction},
			&restrict.Permission{Action: DeleteAction},
		},
	},
	Parents: []string{UserRole},
}

var superAdminRole *restrict.Role = &restrict.Role{
	ID:          SuperAdminRole,
	Description: "SuperAdmin can manage all entities in the system.",
	Grants:      restrict.GrantsMap{},
	Parents:     []string{AdminRole},
}

// Policy - describes Gochat's RBAC policy definition.
var Policy *restrict.PolicyDefinition = &restrict.PolicyDefinition{
	PermissionPresets: restrict.PermissionPresets{
		AccessOwnPreset: &restrict.Permission{
			Conditions: restrict.Conditions{
				&restrict.EqualCondition{
					ID: "isOwner",
					Left: &restrict.ValueDescriptor{
						Source: restrict.ResourceField,
						Field:  "CreatedBy",
					},
					Right: &restrict.ValueDescriptor{
						Source: restrict.SubjectField,
						Field:  "ID",
					},
				},
			},
		},
	},
	Roles: restrict.Roles{
		UserRole:       userRole,
		AdminRole:      adminRole,
		SuperAdminRole: superAdminRole,
	},
}
