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

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(userServiceSuite))
}
