package services

import (
	"github.com/el-Mike/gochat/auth"
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/persist"
	"github.com/el-Mike/gochat/schema"
	"gorm.io/gorm"
)

// AuthService - struct for handling auth related logic.
type AuthService struct {
	broker      *gorm.DB
	userService *UserService
	authManager *auth.AuthManager
}

// NewAuthService - AuthService constructor func.
func NewAuthService() *AuthService {
	return &AuthService{
		broker:      persist.DB,
		userService: NewUserService(),
		authManager: auth.NewAuthManager(),
	}
}

// Login - logs in a user.
// func (as *AuthService) Login(credentials schema.LoginCredentials) (string, error) {

// }

// SignUp - registers a new user, and saves it to DB.
func (as *AuthService) SignUp(credentials schema.SignupPayload) (*models.UserModel, error) {
	hashedPassword, err := as.authManager.HashAndSalt([]byte(credentials.Password))

	if err != nil {
		return nil, err
	}

	userModel := &models.UserModel{
		Password:  hashedPassword,
		Email:     credentials.Email,
		FirstName: credentials.FirstName,
		LastName:  credentials.LastName,
	}

	err = as.userService.SaveUser(userModel)

	if err != nil {
		return nil, err
	}

	return userModel, nil
}
