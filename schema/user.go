package schema

import (
	"github.com/el-Mike/gochat/models"
)

// UserResponse - response for User entity.
type UserResponse struct {
	BaseEntityResponse
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// FromModel - creates UserResponse from UserModel.
func (user *UserResponse) FromModel(model *models.UserModel) error {
	user.ID = model.ID
	user.CreatedAt = model.CreatedAt
	user.UpdatedAt = model.UpdatedAt

	user.Email = model.Email
	user.FirstName = model.FirstName
	user.LastName = model.LastName

	return nil
}
