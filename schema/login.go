package schema

// LoginCredentials - schema for login payload.
type LoginCredentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse - schema for login response.
type LoginResponse struct {
	UserResponse
	Token string `json:"token"`
}
