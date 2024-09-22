package controllers

import (
	"net/http"
	"strconv"

	interfaces "example.com/interfaces/services"
	"github.com/gin-gonic/gin"
)

type RegistrationsController struct {
	registrationService interfaces.IRegistrationService
}

func (controller RegistrationsController) RegisterForEvent(context *gin.Context) {
	eventId, parsingError := strconv.ParseInt(context.Param("id"), 10, 64)

	if parsingError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})
		return
	}

	userId := context.GetInt64("userId")

	err := controller.registrationService.CreateRegistration(eventId, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unexpected error occurred",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "created registration",
	})

}

func (controller RegistrationsController) CancelEventRegistration(context *gin.Context) {
	eventId, parsingError := strconv.ParseInt(context.Param("id"), 10, 64)

	if parsingError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})
		return
	}

	userId := context.GetInt64("userId")

	err := controller.registrationService.DeleteRegistration(eventId, userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unexpected error occurred",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "deleted registration",
	})
}

func NewRegistrationsController(registrationService interfaces.IRegistrationService) *RegistrationsController {
	return &RegistrationsController{
		registrationService: registrationService,
	}
}
