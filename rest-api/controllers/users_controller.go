package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	libInterfaces "example.com/interfaces/lib"
	serviceInterfaces "example.com/interfaces/services"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	userService   serviceInterfaces.IUserService
	jwtAuthorizer libInterfaces.IJwtAuthorizer
}

func (controller UsersController) CreateUser(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payload data",
		})
		return
	}

	err = controller.userService.CreateUser(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Error": fmt.Sprintf("Unable to create user, error: %v\n", err),
		})
		return
	}

	context.JSON(http.StatusCreated, user)
}

func (controller UsersController) Login(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid payload data",
		})
		return
	}

	successfulValidation, err := controller.userService.ValidateCredentials(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Error": fmt.Sprintf("Unable to validate credentials: %v\n", err),
		})
		return
	}

	if !successfulValidation {
		context.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	token, err := controller.jwtAuthorizer.GenerateToken(user.Email, strconv.FormatInt(user.Id, 10))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"Error": fmt.Sprintf("Unable to validate credentials: %v\n", err),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"token":   token,
	})
}

func NewUsersController(userService serviceInterfaces.IUserService, jwtAuthorizer libInterfaces.IJwtAuthorizer) *UsersController {
	return &UsersController{
		userService:   userService,
		jwtAuthorizer: jwtAuthorizer,
	}
}
