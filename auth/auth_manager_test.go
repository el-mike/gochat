package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type AuthManagerMock struct {
	AuthManager
}

func TestNewAuthManager(t *testing.T) {
	am := NewAuthManager()

	assert.NotNil(t, am)
}
