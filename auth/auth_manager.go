package auth

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthManager struct{}

func (am *AuthManager) SignUp() {

}
