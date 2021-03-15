package controllers

import (
	"errors"

	"github.com/el-Mike/gochat/auth"
	"github.com/el-Mike/gochat/common/api"
	"github.com/el-Mike/gochat/common/control"

	"github.com/el-Mike/gochat/schema"
	"github.com/el-Mike/gochat/services"
	"github.com/gin-gonic/gin"
)

// AuthController - struct for handling auth related requests.
type AuthController struct {
	authService *services.AuthService
	userService *services.UserService
	authManager *auth.AuthManager
}

// NewAuthController - AuthController constructor func.
func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
		userService: services.NewUserService(),
		authManager: auth.NewAuthManager(),
	}
}

// Login - authenticates given user and returns a token.
func (ac *AuthController) Login(ctx *gin.Context) (interface{}, *api.APIError) {
	var credentials schema.LoginCredentials

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		return nil, api.NewBadRequestError(err)
	}

	userModel, err := ac.userService.GetUserByEmail(credentials.Email)

	if err != nil {
		return nil, api.NewLoginCredentialsIncorrectError()
	}

	err = ac.authManager.ComparePasswords(userModel.Password, []byte(credentials.Password))

	if err != nil {
		return nil, api.NewLoginCredentialsIncorrectError()
	}

	token, err := ac.authService.Login(userModel)

	if err != nil {
		return nil, api.NewInternalError(err)
	}

	loginResponse := &schema.LoginResponse{}
	err = loginResponse.FromModel(userModel)

	if err != nil {
		return nil, api.NewInternalError(err)
	}

	loginResponse.Token = token

	return loginResponse, nil
}

// Logout - logs out a user.
func (ac *AuthController) Logout(ctx *gin.Context, contextUser *control.ContextUser) (interface{}, *api.APIError) {
	if err := ac.authService.Logout(contextUser); err != nil {
		return nil, api.NewInternalError(err)
	}

	return nil, nil
}

// SignUp - registers a new user
func (ac *AuthController) SignUp(ctx *gin.Context) (interface{}, *api.APIError) {
	var credentials schema.SignupPayload

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		return nil, api.NewBadRequestError(err)
	}

	if _, err := ac.userService.GetUserByEmail(credentials.Email); err == nil {
		return nil, api.NewBadRequestError(errors.New("User already exists."))
	}

	if !schema.ValidatePasswordConfirmation(&credentials) {
		return nil, api.NewBadRequestError(errors.New("Passwords don't match."))
	}

	userModel, err := ac.authService.SignUp(credentials)

	if err != nil {
		return nil, api.NewInternalError(err)
	}

	userResponse := &schema.UserResponse{}
	err = userResponse.FromModel(userModel)

	if err != nil {
		return nil, api.NewInternalError(err)
	}

	return userResponse, nil
}
