package services

import (
	"errors"
	"testing"

	"github.com/el-Mike/gochat/mocks"
	"github.com/el-Mike/gochat/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type userServiceSuite struct {
	suite.Suite
	userService *UserService
	testUserID  uuid.UUID
	testEmail   string
	testRole    string
	testUser    *models.UserModel
}

func (s *userServiceSuite) SetupSuite() {
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
}

func (s *userServiceSuite) SetupTest() {
	s.userService = &UserService{
		broker: mocks.NewGormMock(),
	}
}

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(userServiceSuite))
}

func (s *userServiceSuite) TestNewUserService() {
	userService := NewUserService()

	assert.NotNil(s.T(), userService)
}

func (s *userServiceSuite) TestGetUserByID() {
	userService := s.userService

	gormMock := new(mocks.GormMock)
	gormMock.On(
		"First",
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetDefaultDBResponse())

	userService.broker = gormMock

	user, err := userService.GetUserByID(s.testUserID)

	gormMock.AssertNumberOfCalls(s.T(), "First", 1)

	assert.NotNil(s.T(), user)
	assert.Nil(s.T(), err)
}

func (s *userServiceSuite) TestGetUserByID_Error() {
	userService := s.userService

	gormMock := new(mocks.GormMock)
	gormMock.On(
		"First",
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetErrorDBResponse(errors.New("GormError")))

	userService.broker = gormMock

	user, err := userService.GetUserByID(s.testUserID)

	gormMock.AssertNumberOfCalls(s.T(), "First", 1)

	assert.Nil(s.T(), user)
	assert.NotNil(s.T(), err)
}

func (s *userServiceSuite) TestGetUserByEmail() {
	userService := s.userService

	gormMock := new(mocks.GormMock)
	gormMock.On(
		"FirstWhere",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetDefaultDBResponse())

	userService.broker = gormMock

	user, err := userService.GetUserByEmail(s.testUser.Email)

	gormMock.AssertNumberOfCalls(s.T(), "FirstWhere", 1)

	assert.NotNil(s.T(), user)
	assert.Nil(s.T(), err)
}

func (s *userServiceSuite) TestGetUserByEmail_Error() {
	userService := s.userService

	gormMock := new(mocks.GormMock)
	gormMock.On(
		"FirstWhere",
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetErrorDBResponse(errors.New("GormError")))

	userService.broker = gormMock

	user, err := userService.GetUserByEmail(s.testUser.Email)

	gormMock.AssertNumberOfCalls(s.T(), "FirstWhere", 1)

	assert.Nil(s.T(), user)
	assert.NotNil(s.T(), err)
}

func (s *userServiceSuite) TestGetUsers() {
	userService := s.userService

	gormMock := new(mocks.GormMock)
	gormMock.On(
		"Find",
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetDefaultDBResponse())

	userService.broker = gormMock

	_, err := userService.GetUsers()

	gormMock.AssertNumberOfCalls(s.T(), "Find", 1)

	assert.Nil(s.T(), err)
}

func (s *userServiceSuite) TestGetUsers_Error() {
	userService := s.userService

	gormMock := new(mocks.GormMock)
	gormMock.On(
		"Find",
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetErrorDBResponse(errors.New("GormError")))

	userService.broker = gormMock

	user, err := userService.GetUsers()

	gormMock.AssertNumberOfCalls(s.T(), "Find", 1)

	assert.Nil(s.T(), user)
	assert.NotNil(s.T(), err)
}

func (s *userServiceSuite) TestSaveUser() {
	userService := s.userService

	gormMock := new(mocks.GormMock)
	gormMock.On(
		"Save",
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetDefaultDBResponse())

	userService.broker = gormMock

	err := userService.SaveUser(s.testUser)

	gormMock.AssertNumberOfCalls(s.T(), "Save", 1)
	gormMock.AssertCalled(
		s.T(),
		"Save",
		s.testUser,
	)

	assert.Nil(s.T(), err)
}

func (s *userServiceSuite) TestSaveUser_Error() {
	userService := s.userService

	gormMock := new(mocks.GormMock)
	gormMock.On(
		"Save",
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetErrorDBResponse(errors.New("GormError")))

	userService.broker = gormMock

	err := userService.SaveUser(s.testUser)

	gormMock.AssertNumberOfCalls(s.T(), "Save", 1)
	gormMock.AssertCalled(
		s.T(),
		"Save",
		s.testUser,
	)

	assert.NotNil(s.T(), err)
}

func (s *userServiceSuite) TestDeleteUserByID() {
	userService := s.userService

	gormMock := new(mocks.GormMock)
	gormMock.On(
		"DeleteByID",
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetDefaultDBResponse())

	userService.broker = gormMock

	err := userService.DeleteUserByID(s.testUserID)

	gormMock.AssertNumberOfCalls(s.T(), "DeleteByID", 1)
	gormMock.AssertCalled(
		s.T(),
		"DeleteByID",
		mock.Anything,
		s.testUserID,
	)

	assert.Nil(s.T(), err)
}

func (s *userServiceSuite) TestDeleteUserByID_Error() {
	userService := s.userService

	gormMock := new(mocks.GormMock)
	gormMock.On(
		"DeleteByID",
		mock.Anything,
		mock.Anything,
	).Return(mocks.GetErrorDBResponse(errors.New("GormError")))

	userService.broker = gormMock

	err := userService.DeleteUserByID(s.testUserID)

	gormMock.AssertNumberOfCalls(s.T(), "DeleteByID", 1)
	gormMock.AssertCalled(
		s.T(),
		"DeleteByID",
		mock.Anything,
		s.testUserID,
	)

	assert.NotNil(s.T(), err)
}
