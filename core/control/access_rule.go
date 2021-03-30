package control

import (
	"github.com/el-Mike/gochat/core/control/rbac"
)

// AccessRule - defines a rule that can be used to
// test one's access to a given route.
type AccessRule struct {
	ResourceID string
	Action     rbac.Action
}
