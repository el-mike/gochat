package schema

// SignupPayload - schema for signup payload
type SignupPayload struct {
	Email             string `json:"email" binding:"required,email"`
	Password          string `json:"password" binding:"required,min=8,max=32"`
	ConfirmedPassword string `json:"confirmedPassword" binding:"required,min=8,max=32"`
	FirstName         string `json:"firstName" binding:"required,max=255"`
	LastName          string `json:"lastName" binding:"required,max=255"`
}

// ValidatePasswordConfirmation - returns true when Password and ConfirmedPassword are equal,
// false otherwise.
func ValidatePasswordConfirmation(signupPayload *SignupPayload) bool {
	if signupPayload == nil ||
		signupPayload.Password == "" ||
		signupPayload.ConfirmedPassword == "" {
		return false
	}

	return signupPayload.Password == signupPayload.ConfirmedPassword
}
