package auth

import (
	"errors"
	"fmt"

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

type tokenCreator interface {
	CreateToken(claims jwt.Claims, secret string) (string, error)
	ParseToken(tokenString string, validateFunc jwt.Keyfunc) (*jwt.Token, error)
}

type tokenProvider struct{}

func (tp *tokenProvider) CreateToken(claims jwt.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (tp *tokenProvider) ParseToken(tokenString string, validateFunc jwt.Keyfunc) (*jwt.Token, error) {
	return jwt.Parse(tokenString, validateFunc)
}

// JWTManager - manages operations specific to JWT handling.
type JWTManager struct {
	tokenProvider tokenCreator
}

// NewJWTManager - JWTManager constructor func.
func NewJWTManager() *JWTManager {
	return &JWTManager{
		tokenProvider: &tokenProvider{},
	}
}

// CreateToken - creates a new token for the given user.
func (jm *JWTManager) CreateToken(authUUID, userID, email, role, secret string, time int64) (string, error) {
	if authUUID == "" ||
		userID == "" ||
		email == "" ||
		role == "" ||
		secret == "" ||
		time == 0 {
		return "", errors.New("Missing claims for JWT Token!")
	}

	claims := &JWTClaims{}

	claims.AuthUUID = authUUID
	claims.UserID = userID
	claims.Email = email
	claims.Role = role

	claims.ExpiresAt = time

	return jm.tokenProvider.CreateToken(claims, secret)
}

// ParseToken - parses given token string and returns an instance of jwt.Token.
func (jm *JWTManager) ParseToken(tokenString string, apiSecret string) (*jwt.Token, error) {
	token, err := jm.tokenProvider.ParseToken(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(apiSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
