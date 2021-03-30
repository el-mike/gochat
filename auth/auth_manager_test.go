package auth

import (
	"context"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/el-Mike/gochat/mocks"
	"github.com/el-Mike/gochat/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type jwtManagerMock struct {
	mock.Mock
}

func (jm *jwtManagerMock) CreateToken(authUUID, userID, email, role, apiSecret string, time int64) (string, error) {
	args := jm.Called(authUUID, userID, email, role, apiSecret, time)

	return args.String(0), args.Error(1)
}

func (jm *jwtManagerMock) ParseToken(tokenString string, apiSecret string) (*jwt.Token, error) {
	args := jm.Called(tokenString, apiSecret)

	return args.Get(0).(*jwt.Token), args.Error(1)
}

type cryptoMock struct{}

func (cr *cryptoMock) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	return []byte("test_password"), nil
}

func (cr *cryptoMock) CompareHashAndPassword(hashedPassword, password []byte) error {
	return nil
}

type authManagerSuite struct {
	suite.Suite
	authManager     *AuthManager
	testAuthUUID    string
	testUserID      string
	testEmail       string
	testRole        string
	testSecret      string
	testTime        int64
	testTokenString string
	testUser        *models.UserModel
}

func (s *authManagerSuite) SetupSuite() {
	s.testAuthUUID = "testAuthUUID"
	s.testUserID = "testUserID"
	s.testEmail = "testEmail"
	s.testRole = "testRole"
	s.testSecret = "testSecret"
	s.testTime = time.Now().Unix()
	s.testTokenString = "some_very_long_test_token_string"
	s.testUser = &models.UserModel{}
}

func (s *authManagerSuite) SetupTest() {
	s.authManager = &AuthManager{
		redis:  mocks.NewRedisCacheMock(),
		jwt:    &jwtManagerMock{},
		crypto: &cryptoMock{},
		ctx:    context.Background(),
	}
}

func (s *authManagerSuite) TestLogin() {
	authManager := s.authManager
	jwtMock := new(jwtManagerMock)

	jwtMock.On(
		"CreateToken",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(s.testTokenString, nil)

	authManager.jwt = jwtMock

	token, err := authManager.Login(s.testUser, s.testSecret)

	jwtMock.AssertNumberOfCalls(s.T(), "CreateToken", 1)

	assert.NotEmpty(s.T(), token)
	assert.Equal(s.T(), s.testTokenString, token)
	assert.Nil(s.T(), err)
}

func TestAuthManagerSuite(t *testing.T) {
	suite.Run(t, new(authManagerSuite))
}
