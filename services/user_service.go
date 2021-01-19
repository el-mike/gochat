package services

import (
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/persist"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserService - struct for handling User related logic.
type UserService struct {
	broker *gorm.DB
}

// NewUserService - UserService constructor func.
func NewUserService() *UserService {
	return &UserService{
		broker: persist.DB,
	}
}

// GetUserByID = returns single User with given ID.
func (us *UserService) GetUserByID(id uuid.UUID) (*models.UserModel, error) {
	model := &models.UserModel{}

	err := us.broker.First(&model, id).Error

	return model, err
}

// GetUserByEmail - returns single User with given email.
func (us *UserService) GetUserByEmail(email string) (*models.UserModel, error) {
	model := &models.UserModel{}

	err := us.broker.Where(&models.UserModel{Email: email}).First(model).Error

	return model, err
}

// GetUsers - returns collection of users from DB.
func (us *UserService) GetUsers(users *[]models.UserModel) error {
	return us.broker.Find(users).Error
}

// SaveUser - save single user to DB.
func (us *UserService) SaveUser(user *models.UserModel) error {
	return us.broker.Save(user).Error
}
