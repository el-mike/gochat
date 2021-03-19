package auth_test

import (
	"testing"
	"time"

	"github.com/el-Mike/gochat/auth"
	"github.com/stretchr/testify/assert"
)

var testAuthUUID = "testAuthUUID"
var testUserID = "testUserID"
var testEmail = "testEmail"
var testRole = "testRole"
var testSecret = "testSecret"
var testTime = time.Now().Unix()

func TestCreateTokenValid(t *testing.T) {
	token, err := auth.CreateToken(
		testAuthUUID,
		testUserID,
		testEmail,
		testRole,
		testSecret,
		testTime,
	)

	assert.NotEmpty(t, token)
	assert.Nil(t, err)
}

func TestCreateTokenMissingArgs(t *testing.T) {
	token, err := auth.CreateToken(
		testAuthUUID,
		testUserID,
		testEmail,
		testRole,
		"",
		testTime,
	)

	assert.Empty(t, token)
	assert.NotNil(t, err)
}
