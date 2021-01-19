package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTClaims - JWT Claims for user authentication.
type JWTClaims struct {
	jwt.StandardClaims
	Email    string `json:"email"`
	UserID   string `json:"userID"`
	AuthUUID string `json:"authUUID"`
}

// CreateToken - creates a new token for the given user.
func CreateToken(authUUID, userID, email, secret string, time int64) (string, error) {
	claims := &JWTClaims{}

	claims.AuthUUID = authUUID
	claims.UserID = userID
	claims.Email = email

	claims.ExpiresAt = time

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
