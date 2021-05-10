package controllers

import (
	"errors"

	"github.com/el-Mike/gochat/core/api"
	"github.com/el-Mike/gochat/core/control"
	"github.com/el-Mike/gochat/models"
	"github.com/el-Mike/gochat/schema"
	"github.com/el-Mike/gochat/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
func (uc *UserController) GetMe(ctx *gin.Context, contextUser *control.ContextUser) (interface{}, *api.APIError) {
	id := contextUser.ID

	userModel, err := uc.userService.GetUserByID(id)

	if err != nil {
		return nil, api.NewNotFoundError(models.USER_RESOURCE)
	}

	userResponse := schema.UserResponse{}

	if err := userResponse.FromModel(userModel); err != nil {
		return nil, api.NewInternalError(err)
	}

	return userResponse, nil
}

// GetUsers - returns all the users.
func (uc *UserController) GetUsers(ctx *gin.Context, contextUser *control.ContextUser) (interface{}, *api.APIError) {
	users, err := uc.userService.GetUsers()

	if err != nil {
		return nil, api.NewInternalError(err)
	}

	var userResponses []schema.UserResponse

	for _, userModel := range users {
		userResponse := schema.UserResponse{}

		if err := userResponse.FromModel(userModel); err != nil {
			return nil, api.NewInternalError(err)
		}

		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

// SaveUser - saves single User to DB.
func (uc *UserController) SaveUser(ctx *gin.Context, contextUser *control.ContextUser) (interface{}, *api.APIError) {
	var user models.UserModel

	if err := ctx.ShouldBindJSON(&user); err != nil {
		return nil, api.NewBadRequestError(err)
	}

	existingUser, err := uc.userService.GetUserByEmail(user.Email)
	if err != nil {
		return nil, api.NewInternalError(err)
	}

	if existingUser != nil {
		return nil, api.NewBadRequestError(errors.New("User already exists."))
	}

	if err := uc.userService.SaveUser(&user); err != nil {
		return nil, api.NewInternalError(err)
	}

	return user, nil
}

// DeleteUser - deletes a User with given ID.
func (uc *UserController) DeleteUser(ctx *gin.Context, contextUser *control.ContextUser) (interface{}, *api.APIError) {
	paramId := ctx.Param("id")

	targetId, err := uuid.Parse(paramId)

	if targetId == uuid.Nil || err != nil {
		return nil, api.NewBadRequestError(errors.New("User ID is missing or malformed."))
	}

	userId := contextUser.ID

	if targetId == userId {
		return nil, api.NewBadRequestError(errors.New("You cannot delete yourself."))
	}

	if err := uc.userService.DeleteUserByID(targetId); err != nil {
		return nil, api.NewInternalError(err)
	}

	return nil, nil
}
