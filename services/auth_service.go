package services

import (
	"os"

	"github.com/el-Mike/gochat/auth"
	"github.com/el-Mike/gochat/core/control"
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/persist"
	"github.com/el-Mike/gochat/schema"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// AuthService - struct for handling auth related logic.
type AuthService struct {
	broker      *gorm.DB
	redis       *redis.Client
	userService *UserService
	authManager *auth.AuthManager
}

// NewAuthService - AuthService constructor func.
func NewAuthService() *AuthService {
	return &AuthService{
		broker:      persist.DB,
		redis:       persist.RedisClient,
		userService: NewUserService(),
		authManager: auth.NewAuthManager(),
	}
}

// Login - logs in a user.
func (as *AuthService) Login(user *models.UserModel) (string, error) {
	apiSecret := os.Getenv("API_SECRET")

	token, err := as.authManager.Login(user, apiSecret)

	if err != nil {
		return "", err
	}

	return token, nil
}

// Logout - logs out a user.
func (as *AuthService) Logout(userContext *control.ContextUser) error {
	return as.authManager.Logout(userContext.AuthUUID.String())
}

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
