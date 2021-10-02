package auth

import (
	"errors"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type tokenProviderMock struct {
	mock.Mock
}

func (mtp *tokenProviderMock) CreateToken(claims jwt.Claims, secret string) (string, error) {
	args := mtp.Called(claims, secret)

	return args.String(0), args.Error(1)
}

func (mtp *tokenProviderMock) ParseToken(tokenString string, validateFunc jwt.Keyfunc) (*jwt.Token, error) {
	args := mtp.Called(tokenString, validateFunc)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	token := args.Get(0).(*jwt.Token)

	if _, err := validateFunc(token); err != nil {
		return nil, err
	}

	return token, args.Error(1)
}

type jwtManagerSuite struct {
	suite.Suite
	jwtManager      *JWTManager
	testAuthUUID    string
	testUserID      string
	testEmail       string
	testRole        string
	testSecret      string
	testTime        int64
	testTokenString string
	testToken       *jwt.Token
}

func (s *jwtManagerSuite) SetupSuite() {
	s.testAuthUUID = "testAuthUUID"
	s.testUserID = "testUserID"
	s.testEmail = "testEmail"
	s.testRole = "testRole"
	s.testSecret = "testSecret"
	s.testTime = time.Now().Unix()
	s.testTokenString = "test_token_string"
	s.testToken = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
}

func (s *jwtManagerSuite) SetupTest() {
	s.jwtManager = &JWTManager{
		tokenProvider: &tokenProviderMock{},
	}
}

func TestJWTManagerSuite(t *testing.T) {
	suite.Run(t, new(jwtManagerSuite))
}

func (s *jwtManagerSuite) TestNewJWTManager() {
	jwtManager := NewJWTManager()

	assert.NotNil(s.T(), jwtManager)
	assert.NotNil(s.T(), jwtManager.tokenProvider)
}

func (s *jwtManagerSuite) TestCreateTokenValid() {
	jwtManager := s.jwtManager

	tokenMock := new(tokenProviderMock)
	tokenMock.On(
		"CreateToken",
		mock.Anything,
		mock.Anything,
	).Return(s.testTokenString, nil)

	jwtManager.tokenProvider = tokenMock

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

	tokenMock.AssertNumberOfCalls(s.T(), "CreateToken", 1)
}

func (s *jwtManagerSuite) TestCreateToken_MissingArgs() {
	jwtManager := s.jwtManager

	tokenMock := new(tokenProviderMock)
	tokenMock.On(
		"CreateToken",
		mock.Anything,
		mock.Anything,
	).Return(s.testTokenString, nil)

	jwtManager.tokenProvider = tokenMock

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

func (s *jwtManagerSuite) TestParseToken() {
	jwtManager := s.jwtManager

	tokenMock := new(tokenProviderMock)
	tokenMock.On(
		"ParseToken",
		mock.Anything,
		mock.Anything,
	).Return(s.testToken, nil)

	jwtManager.tokenProvider = tokenMock

	token, err := jwtManager.ParseToken(s.testTokenString, s.testSecret)

	assert.NotNil(s.T(), token)
	assert.Nil(s.T(), err)

	tokenMock.AssertNumberOfCalls(s.T(), "ParseToken", 1)
	tokenMock.AssertCalled(s.T(), "ParseToken", s.testTokenString, mock.Anything)
}

func (s *jwtManagerSuite) TestParseToken_Errors() {
	jwtManager := s.jwtManager

	tokenMock := new(tokenProviderMock)
	tokenMock.On(
		"ParseToken",
		mock.Anything,
		mock.Anything,
	).Return(nil, errors.New("token_error"))

	jwtManager.tokenProvider = tokenMock

	token, err := jwtManager.ParseToken(s.testTokenString, s.testSecret)

	assert.Nil(s.T(), token)
	assert.NotNil(s.T(), err)

	tokenMock.AssertNumberOfCalls(s.T(), "ParseToken", 1)
}

func (s *jwtManagerSuite) TestParseToken_FailingValidation() {
	jwtManager := s.jwtManager

	testToken := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{})

	tokenMock := new(tokenProviderMock)
	tokenMock.On(
		"ParseToken",
		mock.Anything,
		mock.Anything,
	).Return(testToken, nil)

	jwtManager.tokenProvider = tokenMock

	token, err := jwtManager.ParseToken(s.testTokenString, s.testSecret)

	assert.Nil(s.T(), token)
	assert.NotNil(s.T(), err)

	tokenMock.AssertNumberOfCalls(s.T(), "ParseToken", 1)
}
