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

// NewUserController - UserController constructor function
func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

// GetMe - returns user logged in with token sent in request.
func (uc *UserController) GetMe(ctx *gin.Context) {
	contextUser, ok := ctx.Get(api.ContextUserKey)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, api.NewAPIError(500, "User data malformed"))

		return
	}

	id := contextUser.(*api.ContextUser).ID

	userModel, err := uc.userService.GetUserByID(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, api.NewAPIError(400, "User not found"))

		return
	}

	userResponse := schema.UserResponse{}

	userResponse.FromModel(userModel)

	ctx.JSON(http.StatusOK, userResponse)
}

// GetUser - returns single User based on it's ID.
// func (uc *UserController) GetUser(ctx *gin.Context) {

// }

// GetUsers - returns all the users
func (uc *UserController) GetUsers(ctx *gin.Context) {
	var users []models.UserModel

	if err := uc.userService.GetUsers(&users); err != nil {
		ctx.JSON(http.StatusBadRequest, api.FromError(err))

		return
	}

	var userResponses []schema.UserResponse

	for _, userModel := range users {
		userResponse := schema.UserResponse{}

		userResponse.FromModel(&userModel)

		userResponses = append(userResponses, userResponse)
	}

	ctx.JSON(http.StatusOK, userResponses)
}

// SaveUser - saves single User to DB.
func (uc *UserController) SaveUser(ctx *gin.Context) {
	var user models.UserModel

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, api.FromError(err))

		return
	}

	if err := uc.userService.SaveUser(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, api.FromError(err))

		return
	}

	ctx.JSON(http.StatusOK, user)
}
