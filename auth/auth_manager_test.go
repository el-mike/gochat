package auth_test

import (
	"testing"

	"github.com/el-Mike/gochat/auth"
	"github.com/stretchr/testify/assert"
)

type AuthManagerMock struct {
	auth.AuthManager
}

func TestNewAuthManager(t *testing.T) {
	am := auth.NewAuthManager()

	assert.NotNil(t, am)
}
