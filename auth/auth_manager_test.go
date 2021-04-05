package auth

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/el-Mike/gochat/mocks"
	"github.com/el-Mike/gochat/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
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

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*jwt.Token), args.Error(1)
}

type cryptoMock struct {
	mock.Mock
}

func (cr *cryptoMock) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	args := cr.Called(password, cost)

	if _, ok := args.Get(0).([]byte); ok {
		return args.Get(0).([]byte), args.Error(1)
	}

	return []byte(args.String(0)), args.Error(1)
}

func (cr *cryptoMock) CompareHashAndPassword(hashedPassword, password []byte) error {
	args := cr.Called(hashedPassword, password)

	if args.Get(0) == nil {
		return nil
	}

	return args.Error(0)
}

type authManagerSuite struct {
	suite.Suite
	authManager     *AuthManager
	testAuthUUID    uuid.UUID
	testUserID      uuid.UUID
	testEmail       string
	testRole        string
	testSecret      string
	testTime        int64
	testTokenString string
	testToken       *jwt.Token
	testUser        *models.UserModel
	testError       error
	testAuthHeader  string
	testPassword    string
	testHash        string
}

func (s *authManagerSuite) SetupSuite() {
	s.testAuthUUID = uuid.New()
	s.testUserID = uuid.New()
	s.testEmail = "test_email"
	s.testRole = "test_role"
	s.testSecret = "test_secret"
	s.testTime = time.Now().Unix()
	s.testTokenString = "some_very_long_test_token_string"
	s.testToken = jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{})
	s.testUser = &models.UserModel{}
	s.testError = errors.New("test_error")
	s.testAuthHeader = fmt.Sprintf("Bearer %v", s.testTokenString)
	s.testPassword = "test_password"
	s.testHash = "test_hash"
}

func (s *authManagerSuite) SetupTest() {
	s.authManager = &AuthManager{
		cache:  mocks.NewRedisCacheMock(),
		jwt:    &jwtManagerMock{},
		crypto: &cryptoMock{},
		ctx:    context.Background(),
	}
}

func TestAuthManagerSuite(t *testing.T) {
	suite.Run(t, new(authManagerSuite))
}

func (s *authManagerSuite) TestLogin() {
	authManager := s.authManager

	base := models.BaseModel{ID: s.testUserID}
	testUser := &models.UserModel{
		BaseModel: base,
		Email:     s.testEmail,
		Role:      s.testRole,
	}

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

	cacheMock := new(mocks.RedisCacheMock)
	cacheMock.On(
		"Set",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetDefaultCacheResponse())

	authManager.jwt = jwtMock
	authManager.cache = cacheMock

	token, err := authManager.Login(testUser, s.testSecret)

	jwtMock.AssertNumberOfCalls(s.T(), "CreateToken", 1)
	jwtMock.AssertCalled(
		s.T(),
		"CreateToken",
		mock.Anything,
		testUser.ID.String(),
		testUser.Email,
		testUser.Role,
		s.testSecret,
		mock.Anything,
	)

	cacheMock.AssertNumberOfCalls(s.T(), "Set", 1)

	assert.NotEmpty(s.T(), token)
	assert.Equal(s.T(), s.testTokenString, token)
	assert.Nil(s.T(), err)
}

func (s *authManagerSuite) TestLoginWithErrors() {
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
	).Return("", s.testError)

	cacheMock := new(mocks.RedisCacheMock)
	cacheMock.On(
		"Set",
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetErrorCacheResponse(errors.New("cache_error")))

	authManager.jwt = jwtMock
	authManager.cache = cacheMock

	token, err := authManager.Login(s.testUser, s.testSecret)

	assert.Empty(s.T(), token)
	assert.NotNil(s.T(), err)
	jwtMock.AssertNumberOfCalls(s.T(), "CreateToken", 1)
	cacheMock.AssertNumberOfCalls(s.T(), "Set", 0)

	jwtMock = new(jwtManagerMock)
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

	token, err = authManager.Login(s.testUser, s.testSecret)

	assert.Empty(s.T(), token)
	assert.NotNil(s.T(), err)
	jwtMock.AssertNumberOfCalls(s.T(), "CreateToken", 1)
	cacheMock.AssertNumberOfCalls(s.T(), "Set", 1)

}

func (s *authManagerSuite) TestLogout() {
	authManager := s.authManager

	cacheMock := new(mocks.RedisCacheMock)
	cacheMock.On(
		"Del",
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetDefaultCacheResponse())

	authManager.cache = cacheMock

	err := authManager.Logout(s.testAuthUUID.String())

	assert.Nil(s.T(), err)

	cacheMock.AssertCalled(s.T(), "Del", mock.Anything, []string{s.testAuthUUID.String()})
	cacheMock.AssertNumberOfCalls(s.T(), "Del", 1)
}

func (s *authManagerSuite) TestVerifyToken() {
	authManager := s.authManager

	jwtMock := new(jwtManagerMock)
	jwtMock.On(
		"ParseToken",
		s.testTokenString,
		s.testSecret,
	).Return(s.testToken, nil)

	authManager.jwt = jwtMock

	testRequest := mocks.GetTestRequest()
	testRequest.Header.Add("Authorization", s.testAuthHeader)

	token, err := authManager.VerifyToken(testRequest, s.testSecret)

	assert.NotEmpty(s.T(), token)
	assert.Nil(s.T(), err)

	jwtMock.AssertCalled(s.T(), "ParseToken", s.testTokenString, s.testSecret)
	jwtMock.AssertNumberOfCalls(s.T(), "ParseToken", 1)
}

func (s *authManagerSuite) TestVerifyTokenWithMissingToken() {
	authManager := s.authManager

	jwtMock := new(jwtManagerMock)
	jwtMock.On(
		"ParseToken",
		mock.Anything,
		mock.Anything,
	).Return(nil, errors.New("jwt_error"))

	authManager.jwt = jwtMock

	testRequest := mocks.GetTestRequest()
	testRequest.Header.Add("Authorization", "")

	token, err := authManager.VerifyToken(testRequest, s.testSecret)

	assert.Empty(s.T(), token)
	assert.NotNil(s.T(), err)

	jwtMock.AssertCalled(s.T(), "ParseToken", "", s.testSecret)
	jwtMock.AssertNumberOfCalls(s.T(), "ParseToken", 1)
}

func (s *authManagerSuite) TestHashAndSalt() {
	authManager := s.authManager

	cryptoMock := new(cryptoMock)
	cryptoMock.On(
		"GenerateFromPassword",
		mock.Anything,
		mock.Anything,
	).Return(s.testHash, nil)

	authManager.crypto = cryptoMock

	hash, err := authManager.HashAndSalt([]byte(s.testPassword))

	assert.NotEmpty(s.T(), hash)
	assert.IsType(s.T(), "string", hash)
	assert.Nil(s.T(), err)

	cryptoMock.AssertNumberOfCalls(s.T(), "GenerateFromPassword", 1)
	cryptoMock.AssertCalled(s.T(), "GenerateFromPassword", []byte(s.testPassword), bcrypt.MinCost)

}

func (s *authManagerSuite) TestHashAndSaltWithError() {
	authManager := s.authManager

	cryptoMock := new(cryptoMock)
	cryptoMock.On(
		"GenerateFromPassword",
		mock.Anything,
		mock.Anything,
	).Return("", errors.New("crypto_error"))

	authManager.crypto = cryptoMock

	hash, err := authManager.HashAndSalt([]byte(s.testPassword))

	assert.Empty(s.T(), hash)
	assert.NotNil(s.T(), err)

	cryptoMock.AssertNumberOfCalls(s.T(), "GenerateFromPassword", 1)
	cryptoMock.AssertCalled(s.T(), "GenerateFromPassword", []byte(s.testPassword), bcrypt.MinCost)
}

func (s *authManagerSuite) TestComparePasswords() {
	authManager := s.authManager

	cryptoMock := new(cryptoMock)
	cryptoMock.On(
		"CompareHashAndPassword",
		mock.Anything,
		mock.Anything,
	).Return(nil)

	authManager.crypto = cryptoMock

	err := authManager.ComparePasswords(s.testPassword, []byte(s.testPassword))

	assert.Nil(s.T(), err)

	cryptoMock.AssertNumberOfCalls(s.T(), "CompareHashAndPassword", 1)
	cryptoMock.AssertCalled(s.T(), "CompareHashAndPassword", []byte(s.testPassword), []byte(s.testPassword))
}

func (s *authManagerSuite) TestComparePasswordWithError() {
	authManager := s.authManager

	cryptoMock := new(cryptoMock)
	cryptoMock.On(
		"CompareHashAndPassword",
		mock.Anything,
		mock.Anything,
	).Return(errors.New("crypto_error"))

	authManager.crypto = cryptoMock

	err := authManager.ComparePasswords(s.testPassword, []byte(s.testPassword))

	assert.NotNil(s.T(), err)

	cryptoMock.AssertNumberOfCalls(s.T(), "CompareHashAndPassword", 1)
}
