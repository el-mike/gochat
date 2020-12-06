package services

import (
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/persist"
	"gorm.io/gorm"
)

type UserService struct {
	broker *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{
		broker: persist.DB,
	}
}

func (us *UserService) GetUsers(users *[]models.User) error {
	return us.broker.Find(users).Error
}

func (us *UserService) SaveUser(user *models.User) error {
	return us.broker.Save(user).Error
}
