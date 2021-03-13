package models

const USER_MODEL_NAME = "User"

// UserModel - User DB model
type UserModel struct {
	BaseModel
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
