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
	Role     string `json:"role"`
}

// CreateToken - creates a new token for the given user.
func CreateToken(authUUID, userID, email, role, secret string, time int64) (string, error) {
	claims := &JWTClaims{}

	claims.AuthUUID = authUUID
	claims.UserID = userID
	claims.Email = email
	claims.Role = role

	claims.ExpiresAt = time

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}
