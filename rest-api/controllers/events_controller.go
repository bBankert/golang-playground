package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	interfaces "example.com/interfaces/services"
	"example.com/models"
	"github.com/gin-gonic/gin"
)

type EventsController struct {
	eventService interfaces.IEventService
}

func (controller EventsController) GetEvents(context *gin.Context) {

	events, err := controller.eventService.GetEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	context.JSON(http.StatusOK, events)
}

func (controller EventsController) AddEvent(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	event.UserId = context.GetInt64("userId")

	err = controller.eventService.SaveEvent(&event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Created",
		"event":   event,
	})
}

func (controller EventsController) GetEventById(context *gin.Context) {
	eventId, parsingError := strconv.ParseInt(context.Param("id"), 10, 64)

	if parsingError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})
		return
	}

	event, err := controller.eventService.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error trying to fetch event by id, error: %v\n", err),
		})
		return
	}

	//Given that sqlite will auto create ids, if it is 0, then it "does not exist"
	if event.Id == 0 {
		context.JSON(http.StatusNotFound, nil)
		return
	}

	context.JSON(http.StatusOK, event)
}

func (controller EventsController) UpdateEvent(context *gin.Context) {
	eventId, parsingError := strconv.ParseInt(context.Param("id"), 10, 64)

	if parsingError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})
		return
	}
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid event",
		})
	}

	savedEvent, err := controller.eventService.GetEventById(eventId)

	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	requestingUserId := context.GetInt64("userId")

	if savedEvent.UserId != requestingUserId {
		context.Status(http.StatusUnauthorized)
		return
	}

	err = controller.eventService.UpdateEvent(eventId, event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error trying to update event, error: %v\n", err),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Updated",
	})
}

func (controller EventsController) DeleteEvent(context *gin.Context) {
	eventId, parsingError := strconv.ParseInt(context.Param("id"), 10, 64)

	if parsingError != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid event id",
		})
		return
	}

	savedEvent, err := controller.eventService.GetEventById(eventId)

	if err != nil {
		context.Status(http.StatusInternalServerError)
		return
	}

	requestingUserId := context.GetInt64("userId")

	if savedEvent.UserId != requestingUserId {
		context.Status(http.StatusUnauthorized)
		return
	}

	err = controller.eventService.DeleteEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error delete event, error: %v\n", err),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event Deleted",
	})
}

func NewEventsController(eventService interfaces.IEventService) *EventsController {
	return &EventsController{
		eventService: eventService,
	}
}
