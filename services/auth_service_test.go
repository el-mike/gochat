package services

import (
	"errors"
	"os"
	"testing"

	"github.com/el-Mike/gochat/core/control"
	"github.com/el-Mike/gochat/mocks"
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/schema"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userServiceMock struct {
	mock.Mock
}

func (us *userServiceMock) SaveUser(user *models.UserModel) error {
	args := us.Called(user)

	if args.Get(0) == nil {
		return nil
	}

	return args.Error(0)
}

type authManagerMock struct {
	mock.Mock
}

func (am *authManagerMock) Login(user *models.UserModel, apiSecret string) (string, error) {
	args := am.Called(user, apiSecret)

	return args.String(0), args.Error(1)
}

func (am *authManagerMock) Logout(authUUID string) error {
	args := am.Called(authUUID)

	return args.Error(0)
}

func (am *authManagerMock) HashAndSalt(password []byte) (string, error) {
	args := am.Called(password)

	return args.String(0), args.Error(1)
}

type authServiceSuite struct {
	suite.Suite
	authService     *AuthService
	testUserID      uuid.UUID
	testEmail       string
	testRole        string
	testUser        *models.UserModel
	testContextUser *control.ContextUser
	testSecret      string
	testToken       string
	testPassword    string
	testCredentials *schema.SignupPayload
}

func (s *authServiceSuite) SetupSuite() {
	testID := uuid.New()
	testEmail := "test_email@gochat.com"
	testRole := "test_role"

	base := models.BaseModel{ID: testID}

	testUser := &models.UserModel{
		BaseModel: base,
		Email:     testEmail,
		Role:      testRole,
	}

	s.testUserID = testID
	s.testEmail = testEmail
	s.testRole = testRole
	s.testUser = testUser

	s.testContextUser = &control.ContextUser{}

	s.testSecret = "test_api_secret"
	s.testToken = "test_token"

	s.testPassword = "test_password"

	s.testCredentials = &schema.SignupPayload{
		Password:  s.testPassword,
		Email:     s.testEmail,
		FirstName: "first_name",
		LastName:  "last_name",
	}
}

func (s *authServiceSuite) SetupTest() {
	s.authService = &AuthService{
		broker:      mocks.NewGormMock(),
		userService: &userServiceMock{},
		authManager: &authManagerMock{},
	}
}

func TestAuthServiceSuite(t *testing.T) {
	suite.Run(t, new(authServiceSuite))
}

func (s *authServiceSuite) TestLogin() {
	authService := s.authService

	os.Setenv("API_SECRET", s.testSecret)

	authManagerMock := new(authManagerMock)
	authManagerMock.On(
		"Login",
		mock.Anything,
		mock.Anything,
	).Return(s.testToken, nil)

	authService.authManager = authManagerMock

	token, err := authService.Login(s.testUser)

	authManagerMock.AssertNumberOfCalls(s.T(), "Login", 1)

	assert.NotEmpty(s.T(), token)
	assert.Nil(s.T(), err)
}

func (s *authServiceSuite) TestLogin_MissingAPISecret() {
	authService := s.authService

	os.Setenv("API_SECRET", "")

	authManagerMock := new(authManagerMock)
	authManagerMock.On(
		"Login",
		mock.Anything,
		mock.Anything,
	).Return(s.testToken, nil)

	authService.authManager = authManagerMock

	token, err := authService.Login(s.testUser)

	authManagerMock.AssertNumberOfCalls(s.T(), "Login", 0)

	assert.Empty(s.T(), token)
	assert.NotNil(s.T(), err)
}

func (s *authServiceSuite) TestLogin_Error() {
	authService := s.authService

	os.Setenv("API_SECRET", s.testSecret)

	authManagerMock := new(authManagerMock)
	authManagerMock.On(
		"Login",
		mock.Anything,
		mock.Anything,
	).Return("", errors.New("LoginError"))

	authService.authManager = authManagerMock

	token, err := authService.Login(s.testUser)

	authManagerMock.AssertNumberOfCalls(s.T(), "Login", 1)

	assert.Empty(s.T(), token)
	assert.NotNil(s.T(), err)
}

func (s *authServiceSuite) TestLogout() {
	authService := s.authService

	authManagerMock := new(authManagerMock)
	authManagerMock.On(
		"Logout",
		mock.Anything,
	).Return(nil)

	authService.authManager = authManagerMock

	err := authService.Logout(s.testContextUser)

	authManagerMock.AssertNumberOfCalls(s.T(), "Logout", 1)

	assert.Nil(s.T(), err)
}

func (s *authServiceSuite) TestSignUp() {
	authService := s.authService

	authManagerMock := new(authManagerMock)
	authManagerMock.On(
		"HashAndSalt",
		mock.Anything,
	).Return(s.testPassword, nil)

	userServiceMock := new(userServiceMock)
	userServiceMock.On(
		"SaveUser",
		mock.Anything,
	).Return(nil)

	authService.authManager = authManagerMock
	authService.userService = userServiceMock

	user, err := authService.SignUp(*s.testCredentials)

	authManagerMock.AssertNumberOfCalls(s.T(), "HashAndSalt", 1)
	userServiceMock.AssertNumberOfCalls(s.T(), "SaveUser", 1)

	assert.NotNil(s.T(), user)
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), s.testCredentials.Email, user.Email)
	assert.Equal(s.T(), s.testCredentials.FirstName, user.FirstName)
	assert.Equal(s.T(), s.testCredentials.LastName, user.LastName)
}

func (s *authServiceSuite) TestSignUp_PasswordError() {
	authService := s.authService

	authManagerMock := new(authManagerMock)
	authManagerMock.On(
		"HashAndSalt",
		mock.Anything,
	).Return("", errors.New("HashingError"))

	userServiceMock := new(userServiceMock)
	userServiceMock.On(
		"SaveUser",
		mock.Anything,
	).Return(nil)

	authService.authManager = authManagerMock
	authService.userService = userServiceMock

	user, err := authService.SignUp(*s.testCredentials)

	authManagerMock.AssertNumberOfCalls(s.T(), "HashAndSalt", 1)
	userServiceMock.AssertNumberOfCalls(s.T(), "SaveUser", 0)

	assert.Nil(s.T(), user)
	assert.NotNil(s.T(), err)
}

func (s *authServiceSuite) TestSignUp_SaveUserError() {
	authService := s.authService

	authManagerMock := new(authManagerMock)
	authManagerMock.On(
		"HashAndSalt",
		mock.Anything,
	).Return(s.testPassword, nil)

	userServiceMock := new(userServiceMock)
	userServiceMock.On(
		"SaveUser",
		mock.Anything,
	).Return(errors.New("SaveUserError"))

	authService.authManager = authManagerMock
	authService.userService = userServiceMock

	user, err := authService.SignUp(*s.testCredentials)

	authManagerMock.AssertNumberOfCalls(s.T(), "HashAndSalt", 1)
	userServiceMock.AssertNumberOfCalls(s.T(), "SaveUser", 1)

	assert.Nil(s.T(), user)
	assert.NotNil(s.T(), err)
}
