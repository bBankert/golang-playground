package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	interfaces "example.com/interfaces/services"
	"example.com/lib"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	userService interfaces.IUserService
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
		context.Status(http.StatusUnauthorized)
		return
	}

	token, err := lib.GenerateToken(user.Email, strconv.FormatInt(user.Id, 10))

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

func NewUsersController(userService interfaces.IUserService) *UsersController {
	return &UsersController{
		userService: userService,
	}
}
