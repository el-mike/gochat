package auth

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type mockTokenProvider struct{}

func (mtp *mockTokenProvider) CreateToken(claims jwt.Claims, secret string) (string, error) {
	return "test_jwt_signed_token", nil
}

func (mtp *mockTokenProvider) ParseToken(tokenString string, validateFunc jwt.Keyfunc) (*jwt.Token, error) {
	return jwt.NewWithClaims(jwt.SigningMethodNone, &jwt.MapClaims{}), nil
}

type jwtManagerSuite struct {
	suite.Suite
	jwtManager   *JWTManager
	testAuthUUID string
	testUserID   string
	testEmail    string
	testRole     string
	testSecret   string
	testTime     int64
}

func (s *jwtManagerSuite) SetupSuite() {
	s.testAuthUUID = "testAuthUUID"
	s.testUserID = "testUserID"
	s.testEmail = "testEmail"
	s.testRole = "testRole"
	s.testSecret = "testSecret"
	s.testTime = time.Now().Unix()
}

func (s *jwtManagerSuite) SetupTest() {
	s.jwtManager = &JWTManager{
		tokenProvider: &mockTokenProvider{},
	}
}

func (s *jwtManagerSuite) TestCreateTokenValid() {
	jwtManager := s.jwtManager

	token, err := jwtManager.CreateToken(
		s.testAuthUUID,
		s.testUserID,
		s.testEmail,
		s.testRole,
		s.testSecret,
		s.testTime,
	)

	assert.NotEmpty(s.T(), token)
	assert.Nil(s.T(), err)
}

func (s *jwtManagerSuite) TestCreateTokenMissingArgs() {
	jwtManager := s.jwtManager

	token, err := jwtManager.CreateToken(
		s.testAuthUUID,
		s.testUserID,
		s.testEmail,
		s.testRole,
		"",
		s.testTime,
	)

	assert.Empty(s.T(), token)
	assert.NotNil(s.T(), err)
}

func TestJWTManagerSuite(t *testing.T) {
	suite.Run(t, new(jwtManagerSuite))
}
