package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthManager(t *testing.T) {
	am := NewAuthManager()

	assert.NotNil(t, am)
}
