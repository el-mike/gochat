package services

import (
	"context"
	"time"

	"github.com/el-Mike/gochat/auth"
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/persist"
	"github.com/el-Mike/gochat/schema"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AuthService - struct for handling auth related logic.
type AuthService struct {
	broker      *gorm.DB
	redis       *redis.Client
	userService *UserService
	authManager *auth.AuthManager
}

var ctx = context.Background()

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
	authUUID := uuid.New().String()
	userID := user.ID.String()
	email := user.Email
	expiresAt := time.Now().Add(time.Minute * 15).Unix()

	token, err := auth.CreateToken(authUUID, userID, email, expiresAt)

	if err != nil {
		return "", err
	}

	// Saving authorization allows us to double check the token - when user logs out,
	// token will be removed, and no one will be able to use it anymore, even if it's not
	// expired.
	err = as.redis.Set(ctx, authUUID, userID, 0).Err()

	if err != nil {
		return "", err
	}

	return token, nil
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
