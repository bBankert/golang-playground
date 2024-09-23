package services

import (
	"errors"
	"testing"

	"example.com/constants"
	"example.com/mocks"
	"example.com/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RegistrationServiceUnitTestSuite struct {
	suite.Suite
	registrationRepositoryMock mocks.IRegistrationRepository
	eventRepositoryMock        mocks.IEventRepository
	service                    *RegistrationService
}

func TestRegistrationServiceUnitTestSuite(t *testing.T) {
	suite.Run(t, &EventServiceUnitTestSuite{})
}

func (suite *RegistrationServiceUnitTestSuite) SetupTest() {
	suite.eventRepositoryMock = mocks.IEventRepository{}
	suite.registrationRepositoryMock = mocks.IRegistrationRepository{}

	suite.service = NewRegistrationService(
		&suite.registrationRepositoryMock,
		&suite.eventRepositoryMock)
}

func (suite *RegistrationServiceUnitTestSuite) TestCreateRegistration_AttemptsToGetEventById() {

	var expectedEventId, expectedUserId int64 = 1, 12

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(nil, errors.New("test"))

	suite.service.CreateRegistration(expectedEventId, expectedUserId)

	suite.eventRepositoryMock.AssertCalled(suite.T(), "GetEventById", expectedEventId)
	suite.eventRepositoryMock.AssertNumberOfCalls(suite.T(), "GetEventById", 1)
}

// When fetching the event returns an error, pass that error up
func (suite *RegistrationServiceUnitTestSuite) TestCreateRegistration_ReturnsAnError() {

	expectedError := errors.New("test")

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(nil, expectedError)

	err := suite.service.CreateRegistration(1, 1)

	suite.NotNil(err)
	suite.Equal(err, expectedError)
}

// When there is no event for the provided id, return an error
func (suite *RegistrationServiceUnitTestSuite) TestCreateRegistrationWhenNoEventFound_ReturnsAnError() {

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(nil, nil)

	err := suite.service.CreateRegistration(1, 1)

	suite.NotNil(err)
	suite.Equal(err.Error(), constants.NO_EVENT_FOR_ID_ERROR)
}

func (suite *RegistrationServiceUnitTestSuite) TestCreateRegistration_AttemptsToCreateARegistration() {

	var expectedEventId, expectedUserId int64 = 1, 12

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(models.Event{Id: 12}, nil)
	suite.registrationRepositoryMock.On("CreateRegistration", mock.Anything, mock.Anything).Return(errors.New("test"))

	suite.service.CreateRegistration(expectedEventId, expectedUserId)

	suite.registrationRepositoryMock.AssertCalled(suite.T(), "CreateRegistration", expectedEventId, expectedUserId)
	suite.registrationRepositoryMock.AssertNumberOfCalls(suite.T(), "CreateRegistration", 1)
}

// When failing to register, pass up the error
func (suite *RegistrationServiceUnitTestSuite) TestCreateRegistrationWhenUnableToCreateARegistration_ReturnsAnError() {

	expectedError := errors.New("test")

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(models.Event{Id: 12}, nil)
	suite.registrationRepositoryMock.On("CreateRegistration", mock.Anything, mock.Anything).Return(expectedError)

	err := suite.service.CreateRegistration(1, 12)

	suite.NotNil(err)
	suite.Equal(err, expectedError)
}

func (suite *RegistrationServiceUnitTestSuite) TestCreateRegistration_ReturnsNil() {

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(models.Event{Id: 12}, nil)
	suite.registrationRepositoryMock.On("CreateRegistration", mock.Anything, mock.Anything).Return(nil)

	err := suite.service.CreateRegistration(1, 12)

	suite.Nil(err)
}

func (suite *RegistrationServiceUnitTestSuite) TestDeleteRegistration_AttemptsToGetEventById() {

	var expectedEventId, expectedUserId int64 = 1, 12

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(nil, errors.New("test"))

	suite.service.DeleteRegistration(expectedEventId, expectedUserId)

	suite.eventRepositoryMock.AssertCalled(suite.T(), "GetEventById", expectedEventId)
	suite.eventRepositoryMock.AssertNumberOfCalls(suite.T(), "GetEventById", 1)
}

// When fetching the event returns an error, pass that error up
func (suite *RegistrationServiceUnitTestSuite) TestDeleteRegistration_ReturnsAnError() {

	expectedError := errors.New("test")

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(nil, expectedError)

	err := suite.service.DeleteRegistration(1, 1)

	suite.NotNil(err)
	suite.Equal(err, expectedError)
}

// When there is no event for the provided id, return an error
func (suite *RegistrationServiceUnitTestSuite) TestDeleteRegistrationWhenNoEventFound_ReturnsAnError() {

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(nil, nil)

	err := suite.service.DeleteRegistration(1, 1)

	suite.NotNil(err)
	suite.Equal(err.Error(), constants.NO_EVENT_FOR_ID_ERROR)
}

func (suite *RegistrationServiceUnitTestSuite) TestDeleteRegistration_AttemptsToDeleteTheRegistration() {

	var expectedEventId, expectedUserId int64 = 1, 12

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(models.Event{Id: 12}, nil)
	suite.registrationRepositoryMock.On("DeleteRegistration", mock.Anything, mock.Anything).Return(errors.New("test"))

	suite.service.DeleteRegistration(expectedEventId, expectedUserId)

	suite.registrationRepositoryMock.AssertCalled(suite.T(), "DeleteRegistration", expectedEventId, expectedUserId)
	suite.registrationRepositoryMock.AssertNumberOfCalls(suite.T(), "DeleteRegistration", 1)
}

// When failing to register, pass up the error
func (suite *RegistrationServiceUnitTestSuite) TestDeleteRegistrationWhenUnableToCreateARegistration_ReturnsAnError() {

	expectedError := errors.New("test")

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(models.Event{Id: 12}, nil)
	suite.registrationRepositoryMock.On("DeleteRegistration", mock.Anything, mock.Anything).Return(expectedError)

	err := suite.service.DeleteRegistration(1, 12)

	suite.NotNil(err)
	suite.Equal(err, expectedError)
}

func (suite *RegistrationServiceUnitTestSuite) TestDeleteRegistration_ReturnsNil() {

	suite.eventRepositoryMock.On("GetEventById", mock.Anything).Return(models.Event{Id: 12}, nil)
	suite.registrationRepositoryMock.On("DeleteRegistration", mock.Anything, mock.Anything).Return(nil)

	err := suite.service.DeleteRegistration(1, 12)

	suite.Nil(err)
}
