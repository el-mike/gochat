package controllers

import (
	"net/http"

	"github.com/el-Mike/gochat/common/api"

	"github.com/el-Mike/gochat/schema"
	"github.com/el-Mike/gochat/services"
	"github.com/gin-gonic/gin"
)

// AuthController - struct for handling auth related requests.
type AuthController struct {
	authService *services.AuthService
	userService *services.UserService
}

// NewAuthController - AuthController constructor func
func NewAuthController() *AuthController {
	return &AuthController{
		authService: services.NewAuthService(),
		userService: services.NewUserService(),
	}
}

// SignUp - registers a new user
func (ac *AuthController) SignUp(ctx *gin.Context) {
	var credentials schema.SignupPayload

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, api.FromError(err))

		return
	}

	if _, err := ac.userService.GetUserByEmail(credentials.Email); err == nil {
		ctx.JSON(http.StatusBadRequest, api.NewAPIError(400, "User already exists"))

		return
	}

	if schema.ValidatePasswordConfirmation(&credentials) == false {
		ctx.JSON(http.StatusBadRequest, api.NewAPIError(400, "Passwords don't match"))

		return
	}

	userModel, err := ac.authService.SignUp(credentials)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, api.FromError(err))

		return
	}

	userResponse := &schema.UserResponse{}

	err = userResponse.FromModel(userModel)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, api.FromError(err))

		return
	}

	ctx.JSON(http.StatusCreated, userResponse)
}
