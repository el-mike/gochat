package services

import (
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/persist"
	"github.com/google/uuid"
)

// UserService - struct for handling User related logic.
type UserService struct {
	broker persist.DBBroker
}

// NewUserService - UserService constructor func.
func NewUserService() *UserService {
	return &UserService{
		broker: persist.GormBroker,
	}
}

// GetUserByID = returns single User with given ID.
func (us *UserService) GetUserByID(id uuid.UUID) (*models.UserModel, error) {
	model := &models.UserModel{}

	err := us.broker.First(&model, id).Err()
	if err != nil {
		return nil, err
	}

	return model, nil
}

// GetUserByEmail - returns single User with given email.
func (us *UserService) GetUserByEmail(email string) (*models.UserModel, error) {
	model := &models.UserModel{}

	err := us.broker.FirstWhere(model, &models.UserModel{Email: email}).Err()
	if err != nil {
		return nil, err
	}

	return model, nil
}

// GetUsers - returns collection of users from DB.
func (us *UserService) GetUsers() ([]*models.UserModel, error) {
	var users []*models.UserModel

	err := us.broker.Find(&users).Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// SaveUser - save single user to DB.
func (us *UserService) SaveUser(user *models.UserModel) error {
	return us.broker.Save(user).Err()
}
