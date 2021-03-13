package controllers

import (
	"errors"
	"net/http"

	"github.com/el-Mike/gochat/auth"
	"github.com/el-Mike/gochat/common/api"

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

// NewAuthController - AuthController constructor func
func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
		userService: services.NewUserService(),
		authManager: auth.NewAuthManager(),
	}
}

// Login - authenticates given user and returns a token.
func (ac *AuthController) Login(ctx *gin.Context) {
	var credentials schema.LoginCredentials

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(api.ResponseFromError(api.NewBadRequestError(err)))

		return
	}

	userModel, err := ac.userService.GetUserByEmail(credentials.Email)

	if err != nil {
		ctx.JSON(api.ResponseFromError(api.NewLoginCredentialsIncorrectError()))

		return
	}

	err = ac.authManager.ComparePasswords(userModel.Password, []byte(credentials.Password))

	if err != nil {
		ctx.JSON(api.ResponseFromError(api.NewLoginCredentialsIncorrectError()))

		return
	}

	token, err := ac.authService.Login(userModel)

	if err != nil {
		ctx.JSON(api.ResponseFromError(api.NewInternalError(err)))

		return
	}

	loginResponse := &schema.LoginResponse{}
	err = loginResponse.FromModel(userModel)

	if err != nil {
		ctx.JSON(api.ResponseFromError(api.NewInternalError(err)))

		return
	}

	loginResponse.Token = token

	ctx.JSON(http.StatusOK, loginResponse)
}

// SignUp - registers a new user
func (ac *AuthController) SignUp(ctx *gin.Context) {
	var credentials schema.SignupPayload

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(api.ResponseFromError(api.NewBadRequestError(err)))

		return
	}

	if _, err := ac.userService.GetUserByEmail(credentials.Email); err == nil {
		ctx.JSON(api.ResponseFromError(api.NewBadRequestError(errors.New("User already exists."))))

		return
	}

	if schema.ValidatePasswordConfirmation(&credentials) == false {
		ctx.JSON(api.ResponseFromError(api.NewBadRequestError(errors.New("Passwords don't match."))))

		return
	}

	userModel, err := ac.authService.SignUp(credentials)

	if err != nil {
		ctx.JSON(api.ResponseFromError(api.NewInternalError(err)))

		return
	}

	userResponse := &schema.UserResponse{}
	err = userResponse.FromModel(userModel)

	if err != nil {
		ctx.JSON(api.ResponseFromError(api.NewInternalError(err)))

		return
	}

	ctx.JSON(http.StatusCreated, userResponse)
}
