package services

import (
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/persist"
	"github.com/el-Mike/gochat/schema"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthService - struct for handling auth related logic.
type AuthService struct {
	broker      *gorm.DB
	userService *UserService
}

// NewAuthService - AuthService constructor func.
func NewAuthService() *AuthService {
	return &AuthService{
		broker:      persist.DB,
		userService: NewUserService(),
	}
}

// SignUp - registers a new user, and saves it to DB.
func (as *AuthService) SignUp(credentials schema.SignupPayload) (*models.UserModel, error) {
	hashedPassword, err := hashAndSalt([]byte(credentials.Password))

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

func comparePasswords(hashedPassword string, plainPassword []byte) error {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	if err != nil {
		return err
	}

	return nil
}

func hashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
