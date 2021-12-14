package control

import (
	"github.com/el-Mike/restrict"
	"github.com/gin-gonic/gin"
)

// AccessRule - defines a rule that can be used to
// test one's access to a given route.
type AccessRule struct {
	ResourceID       string
	ResourceProvider func(ctx *gin.Context, user *ContextUser) restrict.Resource
	Action           string
}
