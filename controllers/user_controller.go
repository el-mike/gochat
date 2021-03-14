package controllers

import (
	"net/http"

	"github.com/el-Mike/gochat/common/api"
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/schema"
	"github.com/el-Mike/gochat/services"
	"github.com/gin-gonic/gin"
)

// UserController - struct for handling Users related requests.
type UserController struct {
	userService *services.UserService
}

// NewUserController - UserController constructor function.
func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// GetMe - returns user logged in with token sent in request.
func (uc *UserController) GetMe(ctx *gin.Context) {
	contextUser := ctx.MustGet(api.ContextUserKey).(*api.ContextUser)

	id := contextUser.ID

	userModel, err := uc.userService.GetUserByID(id)

	if err != nil {
		ctx.JSON(api.ResponseFromError(api.NewNotFoundError(models.USER_MODEL_NAME)))

		return
	}

	userResponse := schema.UserResponse{}

	userResponse.FromModel(userModel)

	ctx.JSON(http.StatusOK, userResponse)
}

// GetUsers - returns all the users.
func (uc *UserController) GetUsers(ctx *gin.Context) {
	users, err := uc.userService.GetUsers()

	if err != nil {
		ctx.JSON(api.ResponseFromError(api.NewInternalError(err)))

		return
	}

	var userResponses []schema.UserResponse

	for _, userModel := range users {
		userResponse := schema.UserResponse{}

		userResponse.FromModel(userModel)

		userResponses = append(userResponses, userResponse)
	}

	ctx.JSON(http.StatusOK, userResponses)
}

// SaveUser - saves single User to DB.
func (uc *UserController) SaveUser(ctx *gin.Context) {
	var user models.UserModel

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(api.ResponseFromError(api.NewBadRequestError(err)))

		return
	}

	if err := uc.userService.SaveUser(&user); err != nil {
		ctx.JSON(api.ResponseFromError(api.NewInternalError(err)))

		return
	}

	ctx.JSON(http.StatusOK, user)
}
