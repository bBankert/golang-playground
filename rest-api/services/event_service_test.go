package services

import (
	"errors"
	"fmt"
	"testing"

	"example.com/mocks"
	"example.com/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EventServiceUnitTestSuite struct {
	suite.Suite
	eventRepositoryMock mocks.IEventRepository
	service             *EventService
}

func TestEventServiceUnitTestSuite(t *testing.T) {
	suite.Run(t, &EventServiceUnitTestSuite{})
}

func (suite *EventServiceUnitTestSuite) SetupTest() {
	suite.eventRepositoryMock = mocks.IEventRepository{}

	suite.service = NewEventService(&suite.eventRepositoryMock)
}

func (suite *EventServiceUnitTestSuite) TestSaveEvent_AttemptToCreateAnEvent() {

	suite.eventRepositoryMock.On("AddEvent", mock.Anything).Return(errors.New("test"))

	suite.service.SaveEvent(&models.Event{})

	suite.eventRepositoryMock.AssertCalled(suite.T(), "AddEvent", mock.AnythingOfType(fmt.Sprintf("%T", &models.Event{})))
	suite.eventRepositoryMock.AssertNumberOfCalls(suite.T(), "AddEvent", 1)

}

// When an error occurs during db access, return the error
func (suite *EventServiceUnitTestSuite) TestSaveEvent_ReturnsError() {

	mockError := errors.New("test")

	suite.eventRepositoryMock.On("AddEvent", mock.Anything).Return(mockError)

	err := suite.service.SaveEvent(&models.Event{})

	suite.NotNil(err)
	suite.Equal(err.Error(), mockError.Error())

}

func (suite *EventServiceUnitTestSuite) TestSaveEvent_ReturnsNil() {

	suite.eventRepositoryMock.On("AddEvent", mock.Anything).Return(nil)

	result := suite.service.SaveEvent(&models.Event{})

	suite.Nil(result)
}

func (suite *EventServiceUnitTestSuite) TestGetEvents_AttemptToCreateAnEvent() {

	suite.eventRepositoryMock.On("GetEvents").Return(nil, errors.New("test"))

	suite.service.GetEvents()

	suite.eventRepositoryMock.AssertCalled(suite.T(), "GetEvents")
	suite.eventRepositoryMock.AssertNumberOfCalls(suite.T(), "GetEvents", 1)

}

// When an error occurs during db access, return the error
func (suite *EventServiceUnitTestSuite) TestGetEvents_ReturnsError() {

	mockError := errors.New("test")

	suite.eventRepositoryMock.On("GetEvents").Return(nil, mockError)

	_, err := suite.service.GetEvents()

	suite.NotNil(err)
	suite.Equal(err.Error(), mockError.Error())

}

func (suite *EventServiceUnitTestSuite) TestGetEvents_ReturnsExistingEvents() {

	var mockEvents = []models.Event{
		{
			Id:       1,
			Name:     "some name",
			Location: "some location",
		},
	}

	suite.eventRepositoryMock.On("GetEvents", mock.Anything).Return(mockEvents, nil)

	result, _ := suite.service.GetEvents()

	suite.NotNil(result)
	suite.Equal(result, mockEvents)
}

func (suite *EventServiceUnitTestSuite) TestGetEventById_AttemptsToFetchEventById() {

	var expectedEventId int64 = 1

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(nil, errors.New("test"))

	suite.service.GetEventById(expectedEventId)

	suite.eventRepositoryMock.AssertCalled(suite.T(), "GetEventById", expectedEventId)
	suite.eventRepositoryMock.AssertNumberOfCalls(suite.T(), "GetEventById", 1)
}

// When an error occurs during db access, return the error
func (suite *EventServiceUnitTestSuite) TestGetEventById_ReturnsAnError() {

	expectedError := errors.New("test")

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(nil, expectedError)

	_, err := suite.service.GetEventById(1)

	suite.NotNil(err)
	suite.Equal(err, expectedError)
}

func (suite *EventServiceUnitTestSuite) TestGetEventById_ReturnsTheEvent() {

	var mockEvent = models.Event{
		Id:       1,
		Name:     "some name",
		Location: "some location",
	}

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(&mockEvent, nil)

	result, _ := suite.service.GetEventById(1)

	suite.NotNil(result)
	suite.Equal(*result, mockEvent)
}

func (suite *EventServiceUnitTestSuite) TestUpdateEvent_AttemptsToUpdateTheEvent() {

	var expectedEventId int64 = 1

	var expectedEvent = models.Event{
		Id:       1,
		Name:     "some name",
		Location: "some location",
	}

	suite.eventRepositoryMock.On("UpdateEvent", mock.Anything, mock.Anything).Return(errors.New("test"))

	suite.service.UpdateEvent(expectedEventId, expectedEvent)

	suite.eventRepositoryMock.AssertCalled(suite.T(), "UpdateEvent", expectedEventId, expectedEvent)
	suite.eventRepositoryMock.AssertNumberOfCalls(suite.T(), "UpdateEvent", 1)
}

// When an error occurs during db access, return the error
func (suite *EventServiceUnitTestSuite) TestUpdateEvent_ReturnsAnError() {

	expectedError := errors.New("test")

	suite.eventRepositoryMock.On("UpdateEvent", mock.Anything, mock.Anything).Return(expectedError)

	err := suite.service.UpdateEvent(int64(1), models.Event{})

	suite.NotNil(err)
	suite.Equal(err, expectedError)
}

func (suite *EventServiceUnitTestSuite) TestUpdateEvent_ReturnsNil() {

	suite.eventRepositoryMock.On("UpdateEvent", mock.Anything, mock.Anything).Return(nil)

	err := suite.service.UpdateEvent(int64(1), models.Event{})

	suite.Nil(err)
}

func (suite *EventServiceUnitTestSuite) TestDeleteEvent_AttemptsToDeleteTheEvent() {

	var expectedEventId int64 = 1

	suite.eventRepositoryMock.On("DeleteEvent", mock.Anything).Return(errors.New("test"))

	suite.service.DeleteEvent(expectedEventId)

	suite.eventRepositoryMock.AssertCalled(suite.T(), "DeleteEvent", expectedEventId)
	suite.eventRepositoryMock.AssertNumberOfCalls(suite.T(), "DeleteEvent", 1)
}

// When an error occurs during db access, return the error
func (suite *EventServiceUnitTestSuite) TestDeleteEvent_ReturnsAnError() {

	expectedError := errors.New("test")

	suite.eventRepositoryMock.On("DeleteEvent", mock.Anything).Return(expectedError)

	err := suite.service.DeleteEvent(1)

	suite.NotNil(err)
	suite.Equal(err, expectedError)
}

func (suite *EventServiceUnitTestSuite) TestDeleteEvent_ReturnsNil() {

	suite.eventRepositoryMock.On("DeleteEvent", mock.Anything).Return(nil)

	err := suite.service.DeleteEvent(1)

	suite.Nil(err)
}
