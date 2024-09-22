package controllers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"example.com/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RegistrationsControllerUnitTestSuite struct {
	suite.Suite
	mockContext             *gin.Context
	registrationServiceMock mocks.IRegistrationService
	mockResponseWriter      *httptest.ResponseRecorder
	controller              *RegistrationsController
}

func TestRegistrationsControllerUnitTestSuite(t *testing.T) {
	suite.Run(t, &RegistrationsControllerUnitTestSuite{})
}

func (suite *RegistrationsControllerUnitTestSuite) SetupTest() {

	suite.mockResponseWriter = httptest.NewRecorder()

	suite.mockContext, _ = gin.CreateTestContext(suite.mockResponseWriter)

	suite.registrationServiceMock = mocks.IRegistrationService{}

	suite.controller = NewRegistrationsController(&suite.registrationServiceMock)
}

// When the request param is missing, return a bad request
func (suite *RegistrationsControllerUnitTestSuite) TestRegisterForEventWhenParamIsMissing_ReturnsBadRequest() {

	suite.controller.RegisterForEvent(suite.mockContext)

	suite.Equal(http.StatusBadRequest, suite.mockResponseWriter.Code)
}

// When the request param is not a valid id, return a bad request
func (suite *RegistrationsControllerUnitTestSuite) TestRegisterForEventWhenParamIsInvalid_ReturnsBadRequest() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "bar",
		},
	}

	suite.controller.RegisterForEvent(suite.mockContext)

	suite.Equal(http.StatusBadRequest, suite.mockResponseWriter.Code)
}

func (suite *RegistrationsControllerUnitTestSuite) TestRegisterForEvent_CreatesRegistration() {

	var expectedEventId, expectedUserId int64 = 1, 123

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: strconv.FormatInt(expectedEventId, 10),
		},
	}

	suite.mockContext.Set("userId", expectedUserId)

	suite.registrationServiceMock.On("CreateRegistration", mock.Anything, mock.Anything).Return(errors.New("test error"))

	suite.controller.RegisterForEvent(suite.mockContext)

	suite.registrationServiceMock.AssertCalled(suite.T(), "CreateRegistration", expectedEventId, expectedUserId)
}

// When failing to create a registration return internal server error
func (suite *RegistrationsControllerUnitTestSuite) TestRegisterForEvent_ReturnsInternalServerError() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.mockContext.Set("userId", int64(12))

	suite.registrationServiceMock.On("CreateRegistration", mock.Anything, mock.Anything).Return(errors.New("test error"))

	suite.controller.RegisterForEvent(suite.mockContext)

	suite.Equal(http.StatusInternalServerError, suite.mockResponseWriter.Code)
}

func (suite *RegistrationsControllerUnitTestSuite) TestRegisterForEvent_ReturnsCreated() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.mockContext.Set("userId", int64(12))

	suite.registrationServiceMock.On("CreateRegistration", mock.Anything, mock.Anything).Return(nil)

	suite.controller.RegisterForEvent(suite.mockContext)

	suite.Equal(http.StatusCreated, suite.mockResponseWriter.Code)
}

// When the request param is missing, return a bad request
func (suite *RegistrationsControllerUnitTestSuite) TestCancelEventRegistrationWhenParamIsMissing_ReturnsBadRequest() {

	suite.controller.CancelEventRegistration(suite.mockContext)

	suite.Equal(http.StatusBadRequest, suite.mockResponseWriter.Code)
}

// When the request param is not a valid id, return a bad request
func (suite *RegistrationsControllerUnitTestSuite) TestCancelEventRegistrationWhenParamIsInvalid_ReturnsBadRequest() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "bar",
		},
	}

	suite.controller.CancelEventRegistration(suite.mockContext)

	suite.Equal(http.StatusBadRequest, suite.mockResponseWriter.Code)
}

func (suite *RegistrationsControllerUnitTestSuite) TestCancelEventRegistration_DeletesRegistration() {

	var expectedEventId, expectedUserId int64 = 1, 123

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: strconv.FormatInt(expectedEventId, 10),
		},
	}

	suite.mockContext.Set("userId", expectedUserId)

	suite.registrationServiceMock.On("DeleteRegistration", mock.Anything, mock.Anything).Return(errors.New("test error"))

	suite.controller.CancelEventRegistration(suite.mockContext)

	suite.registrationServiceMock.AssertCalled(suite.T(), "DeleteRegistration", expectedEventId, expectedUserId)
}

// When failing to delete a registration return internal server error
func (suite *RegistrationsControllerUnitTestSuite) TestCancelEventRegistration_ReturnsInternalServerError() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.mockContext.Set("userId", int64(12))

	suite.registrationServiceMock.On("DeleteRegistration", mock.Anything, mock.Anything).Return(errors.New("test error"))

	suite.controller.CancelEventRegistration(suite.mockContext)

	suite.Equal(http.StatusInternalServerError, suite.mockResponseWriter.Code)
}

func (suite *RegistrationsControllerUnitTestSuite) TestCancelEventRegistration_ReturnsOk() {

	suite.mockContext.Params = gin.Params{
		{
			Key:   "id",
			Value: "1",
		},
	}

	suite.mockContext.Set("userId", int64(12))

	suite.registrationServiceMock.On("DeleteRegistration", mock.Anything, mock.Anything).Return(nil)

	suite.controller.CancelEventRegistration(suite.mockContext)

	suite.Equal(http.StatusOK, suite.mockResponseWriter.Code)
}
