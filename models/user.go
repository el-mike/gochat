package models

// USER_RESOURCE -name of User resource.
const USER_RESOURCE = "User"

// UserModel - User DB model.
type UserModel struct {
	BaseModel
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}

// ResourceName - returns the name of User resource.
func (um *UserModel) GetResourceName() string {
	return USER_RESOURCE
}
