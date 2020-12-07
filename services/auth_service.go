package services

import (
	"github.com/el-Mike/gochat/auth"
	"github.com/el-Mike/gochat/persist"
	"gorm.io/gorm"
)

type AuthService struct {
	broker *gorm.DB
}

func NewAuthService() *AuthService {
	return &AuthService{
		broker: persist.DB,
	}
}

func (as *AuthService) SignUp(credentials auth.Credentials) {

}
