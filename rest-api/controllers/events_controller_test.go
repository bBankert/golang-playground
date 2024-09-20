package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"example.com/mocks"
	"example.com/models"
	"example.com/test_utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EventsControllerUnitTestSuite struct {
	suite.Suite
	mockContext        *gin.Context
	eventServiceMock   mocks.IEventService
	mockResponseWriter *httptest.ResponseRecorder
	controller         *EventsController
}

func TestEventsControllerUnitTestSuite(t *testing.T) {
	suite.Run(t, &EventsControllerUnitTestSuite{})
}

func (suite *EventsControllerUnitTestSuite) SetupTest() {

	suite.mockResponseWriter = httptest.NewRecorder()

	suite.mockContext, _ = gin.CreateTestContext(suite.mockResponseWriter)

	suite.eventServiceMock = mocks.IEventService{}

	suite.controller = NewEventsController(&suite.eventServiceMock)
}

// Verify an internal server error is returned when an error occurs fetching events
func (suite *EventsControllerUnitTestSuite) TestGetEvents_ReturnsInternalServerError() {

	suite.eventServiceMock.On("GetEvents").Return(nil, errors.New("test error"))

	suite.controller.GetEvents(suite.mockContext)

	suite.Equal(http.StatusInternalServerError, suite.mockResponseWriter.Code)

}

func (suite *EventsControllerUnitTestSuite) TestGetEvents_FetchesEvents() {

	suite.eventServiceMock.On("GetEvents").Return([]models.Event{}, nil)

	suite.controller.GetEvents(suite.mockContext)

	suite.eventServiceMock.AssertCalled(suite.T(), "GetEvents")
	suite.eventServiceMock.AssertNumberOfCalls(suite.T(), "GetEvents", 1)
}

func (suite *EventsControllerUnitTestSuite) TestGetEvents_ReturnsOk() {

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")
	var mockEvent = models.Event{
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
	}

	suite.eventServiceMock.On("GetEvents").Return([]models.Event{
		mockEvent,
	}, nil)

	suite.controller.GetEvents(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusOK, response.StatusCode)

	serializedEvent, _ := json.Marshal([]models.Event{
		mockEvent,
	})
	suite.Equal(response.Body, string(serializedEvent))
}

// When an invalid payload is sent, it should return a bad request
func (suite *EventsControllerUnitTestSuite) TestAddEvents_ReturnsBadRequest() {

	//Missing required bindings name & date
	test_utils.SetRequestBody(models.Event{
		Description: "some description",
		Location:    "some location",
	}, suite.mockContext)

	suite.controller.AddEvent(suite.mockContext)

	suite.Equal(http.StatusBadRequest, suite.mockResponseWriter.Code)

}

func (suite *EventsControllerUnitTestSuite) TestAddEvents_AttemptsToSaveTheNewEvent() {

	var expectedEvent = models.Event{
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        time.Now(),
	}

	test_utils.SetRequestBody(expectedEvent, suite.mockContext)

	suite.eventServiceMock.On("SaveEvent", mock.Anything).Return(nil)

	suite.controller.AddEvent(suite.mockContext)

	//Cannot verify a pointer that would be different given how gin will make a copy of the expected param
	suite.eventServiceMock.AssertCalled(suite.T(), "SaveEvent", mock.AnythingOfType("*models.Event"))
	suite.eventServiceMock.AssertNumberOfCalls(suite.T(), "SaveEvent", 1)
}

// When failing to save the new event, we have to return an error
func (suite *EventsControllerUnitTestSuite) TestAddEvents_ReturnsInternalServerError() {

	var expectedEvent = models.Event{
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        time.Now(),
	}

	test_utils.SetRequestBody(expectedEvent, suite.mockContext)

	suite.eventServiceMock.On("SaveEvent", mock.Anything).Return(errors.New("test"))

	suite.controller.AddEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusInternalServerError)
}

func (suite *EventsControllerUnitTestSuite) TestAddEvents_ReturnsCreated() {

	var expectedEvent = models.Event{
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        time.Now(),
	}

	test_utils.SetRequestBody(expectedEvent, suite.mockContext)

	suite.eventServiceMock.On("SaveEvent", mock.Anything).Return(nil)

	suite.controller.AddEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusCreated)

	data, _ := json.Marshal(expectedEvent)

	suite.Contains(response.Body, string(data))
}

// When there is a malformed or missing id param, it should return bad request
func (suite *EventsControllerUnitTestSuite) TestGetEventByIdMissingParam_ReturnsBadRequest() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "foo",
			Value: "bar",
		},
	}

	suite.controller.GetEventById(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusBadRequest)
}

func (suite *EventsControllerUnitTestSuite) TestGetEventByIdMalformedParam_ReturnsBadRequest() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "bar",
		},
	}

	suite.controller.GetEventById(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusBadRequest)
}

func (suite *EventsControllerUnitTestSuite) TestGetEventById_AttemptsToFetchTheEvent() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(nil, errors.New("test"))

	suite.controller.GetEventById(suite.mockContext)

	//Type is validated as well, so we need to have the right type of int
	suite.eventServiceMock.AssertCalled(suite.T(), "GetEventById", int64(1))
	suite.eventServiceMock.AssertNumberOfCalls(suite.T(), "GetEventById", 1)
}

// When an error is returned trying to fetch the event, returns an internal server error
func (suite *EventsControllerUnitTestSuite) TestGetEventById_ReturnsInternalServerError() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(nil, errors.New("test"))

	suite.controller.GetEventById(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusInternalServerError)
}

// When the default pointer value for an event is returned, we can interpret that as "not found" for
// the event query
func (suite *EventsControllerUnitTestSuite) TestGetEventById_ReturnsNotFound() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(&models.Event{}, nil)

	suite.controller.GetEventById(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusNotFound)
}

func (suite *EventsControllerUnitTestSuite) TestGetEventById_ReturnsOk() {

	var expectedEvent = models.Event{
		Id:   123,
		Name: "some name",
	}

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(&expectedEvent, nil)

	suite.controller.GetEventById(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusOK)

	data, _ := json.Marshal(expectedEvent)

	suite.Contains(string(data), response.Body)
}

// When there is a malformed or missing id param, it should return bad request
func (suite *EventsControllerUnitTestSuite) TestUpdateEventMissingParam_ReturnsBadRequest() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "foo",
			Value: "bar",
		},
	}

	suite.controller.UpdateEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusBadRequest)
}

func (suite *EventsControllerUnitTestSuite) TestUpdateEventMalformedParam_ReturnsBadRequest() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "bar",
		},
	}

	suite.controller.UpdateEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusBadRequest)
}

// When missing required params or missing the body entirely, it should return bad request
func (suite *EventsControllerUnitTestSuite) TestUpdateEventMalformedBody_ReturnsBadRequest() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "bar",
		},
	}

	test_utils.SetRequestBody(models.Event{
		Name: "some name",
	}, suite.mockContext)

	suite.controller.UpdateEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusBadRequest)
}

func (suite *EventsControllerUnitTestSuite) TestUpdateEvent_AttemptsToFetchTheEvent() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(nil, errors.New("test"))

	suite.controller.UpdateEvent(suite.mockContext)

	//Type is validated as well, so we need to have the right type of int
	suite.eventServiceMock.AssertCalled(suite.T(), "GetEventById", int64(1))
	suite.eventServiceMock.AssertNumberOfCalls(suite.T(), "GetEventById", 1)
}

// When the user attempting to update the event, attempts to update the event, it should return
// unauthorized
func (suite *EventsControllerUnitTestSuite) TestUpdateEventNotTheCreator_ReturnsUnauthorized() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	var expectedEvent = models.Event{
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        time.Now(),
	}

	test_utils.SetRequestBody(expectedEvent, suite.mockContext)

	suite.mockContext.Set("userId", int64(1))

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(&models.Event{
		Id:     123,
		UserId: 12,
	}, nil)

	suite.controller.UpdateEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *EventsControllerUnitTestSuite) TestUpdateEvent_AttemptsToUpdateTheEvent() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")

	var expectedEvent = models.Event{
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
	}

	test_utils.SetRequestBody(expectedEvent, suite.mockContext)

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(&models.Event{
		Id:     123,
		UserId: 12,
	}, nil)

	suite.mockContext.Set("userId", int64(12))

	suite.eventServiceMock.On("UpdateEvent", mock.Anything, mock.Anything).Return(errors.New("test"))

	suite.controller.UpdateEvent(suite.mockContext)

	//Type is validated as well, so we need to have the right type of int
	suite.eventServiceMock.AssertCalled(suite.T(), "UpdateEvent", int64(1), expectedEvent)
	suite.eventServiceMock.AssertNumberOfCalls(suite.T(), "UpdateEvent", 1)
}

// When failing to update the event, return an internal server error
func (suite *EventsControllerUnitTestSuite) TestUpdateEvent_ReturnsInternalServerError() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")

	var expectedEvent = models.Event{
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
	}

	test_utils.SetRequestBody(expectedEvent, suite.mockContext)

	suite.mockContext.Set("userId", int64(12))

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(&models.Event{
		Id:     123,
		UserId: 12,
	}, nil)

	suite.eventServiceMock.On("UpdateEvent", mock.Anything, mock.Anything).Return(errors.New("test"))

	suite.controller.UpdateEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusInternalServerError, response.StatusCode)
}

func (suite *EventsControllerUnitTestSuite) TestUpdateEvent_ReturnsOk() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	var expectedEvent = models.Event{
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        time.Now(),
	}

	test_utils.SetRequestBody(expectedEvent, suite.mockContext)

	suite.mockContext.Set("userId", int64(12))

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(&models.Event{
		Id:     123,
		UserId: 12,
	}, nil)

	suite.eventServiceMock.On("UpdateEvent", mock.Anything, mock.Anything).Return(nil)

	suite.controller.UpdateEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusOK, response.StatusCode)
}

// When there is a malformed or missing id param, it should return bad request
func (suite *EventsControllerUnitTestSuite) TestDeleteEventMissingParam_ReturnsBadRequest() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "foo",
			Value: "bar",
		},
	}

	suite.controller.DeleteEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusBadRequest)
}

func (suite *EventsControllerUnitTestSuite) TestDeleteEventMalformedParam_ReturnsBadRequest() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "bar",
		},
	}

	suite.controller.DeleteEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(response.StatusCode, http.StatusBadRequest)
}

func (suite *EventsControllerUnitTestSuite) TestDeleteEvent_AttemptsToFetchTheEvent() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(nil, errors.New("test"))

	suite.controller.DeleteEvent(suite.mockContext)

	//Type is validated as well, so we need to have the right type of int
	suite.eventServiceMock.AssertCalled(suite.T(), "GetEventById", int64(1))
	suite.eventServiceMock.AssertNumberOfCalls(suite.T(), "GetEventById", 1)
}

// When the user is attempting to delete the event, attempts to update the event, it should return
// unauthorized
func (suite *EventsControllerUnitTestSuite) TestDeleteEventNotTheCreator_ReturnsUnauthorized() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.mockContext.Set("userId", int64(1))

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(&models.Event{
		Id:     123,
		UserId: 12,
	}, nil)

	suite.controller.DeleteEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *EventsControllerUnitTestSuite) TestDeleteEvent_AttemptsToDeleteTheEvent() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(&models.Event{
		Id:     123,
		UserId: 12,
	}, nil)

	suite.mockContext.Set("userId", int64(12))

	suite.eventServiceMock.On("DeleteEvent", mock.Anything).Return(errors.New("test"))

	suite.controller.DeleteEvent(suite.mockContext)

	//Type is validated as well, so we need to have the right type of int
	suite.eventServiceMock.AssertCalled(suite.T(), "DeleteEvent", int64(1))
	suite.eventServiceMock.AssertNumberOfCalls(suite.T(), "DeleteEvent", 1)
}

// When failing to update the event, return an internal server error
func (suite *EventsControllerUnitTestSuite) TestDeleteEvent_ReturnsInternalServerError() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.mockContext.Set("userId", int64(12))

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(&models.Event{
		Id:     123,
		UserId: 12,
	}, nil)

	suite.eventServiceMock.On("DeleteEvent", mock.Anything).Return(errors.New("test"))

	suite.controller.DeleteEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusInternalServerError, response.StatusCode)
}

func (suite *EventsControllerUnitTestSuite) TestDeleteEvent_ReturnsOk() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.mockContext.Set("userId", int64(12))

	suite.eventServiceMock.On("GetEventById", mock.Anything).Return(&models.Event{
		Id:     123,
		UserId: 12,
	}, nil)

	suite.eventServiceMock.On("DeleteEvent", mock.Anything).Return(nil)

	suite.controller.DeleteEvent(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusOK, response.StatusCode)
}
