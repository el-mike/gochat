package auth

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Claims - JWT Claims for user authentication.
type Claims struct {
	jwt.StandardClaims
	Email    string `json:"email"`
	UserID   string `json:"userId"`
	AuthUUID string
}

// AuthManager - manages auth related operations.
type AuthManager struct{}

// NewAuthManager - AuthManager constructor func.
func NewAuthManager() *AuthManager {
	return &AuthManager{}
}

// HashAndSalt - returned bcrypt hashed password.
func (am *AuthManager) HashAndSalt(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// ComparePasswords - checks if password matches it's hased version.
func (am *AuthManager) ComparePasswords(hashedPassword string, plainPassword []byte) error {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	if err != nil {
		return err
	}

	return nil
}
