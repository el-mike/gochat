package auth

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

var testAuthUUID = "testAuthUUID"
var testUserID = "testUserID"
var testEmail = "testEmail"
var testRole = "testRole"
var testSecret = "testSecret"
var testTime = time.Now().Unix()

type mockTokenProvider struct{}

func (mtp *mockTokenProvider) CreateToken(claims jwt.Claims, secret string) (string, error) {
	return "test_jwt_signed_token", nil
}

func (mtp *mockTokenProvider) ParseToken(tokenString string, validateFunc jwt.Keyfunc) (*jwt.Token, error) {
	return jwt.NewWithClaims(jwt.SigningMethodNone, &jwt.MapClaims{}), nil
}

var jwtManager = &JWTManager{
	tokenProvider: &mockTokenProvider{},
}

func TestCreateTokenValid(t *testing.T) {
	token, err := jwtManager.CreateToken(
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
	token, err := jwtManager.CreateToken(
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
