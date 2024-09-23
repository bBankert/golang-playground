package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"example.com/mocks"
	"example.com/models"
	"example.com/test_utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UsersControllerUnitTestSuite struct {
	suite.Suite
	mockContext        *gin.Context
	userServiceMock    mocks.IUserService
	jwtAuthorizerMock  mocks.IJwtAuthorizer
	mockResponseWriter *httptest.ResponseRecorder
	controller         *UsersController
}

func TestUsersControllerUnitTestSuite(t *testing.T) {
	suite.Run(t, &UsersControllerUnitTestSuite{})
}

func (suite *UsersControllerUnitTestSuite) SetupTest() {

	suite.mockResponseWriter = httptest.NewRecorder()

	suite.mockContext, _ = gin.CreateTestContext(suite.mockResponseWriter)

	suite.userServiceMock = mocks.IUserService{}
	suite.jwtAuthorizerMock = mocks.IJwtAuthorizer{}

	suite.controller = NewUsersController(&suite.userServiceMock, &suite.jwtAuthorizerMock)
}

// When provided an invalid payload, should return bad request
func (suite *UsersControllerUnitTestSuite) TestCreateUser_ReturnsBadRequest() {

	test_utils.SetRequestBody(models.User{
		Email: "some email",
	}, suite.mockContext)

	suite.controller.CreateUser(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusBadRequest, response.StatusCode)
}

func (suite *UsersControllerUnitTestSuite) TestCreateUser_AttemptsToCreateAUser() {

	var mockUser = models.User{
		Email:    "some email",
		Password: "some password",
	}

	test_utils.SetRequestBody(mockUser, suite.mockContext)

	suite.userServiceMock.On("CreateUser", mock.Anything).Return(errors.New("test"))

	suite.controller.CreateUser(suite.mockContext)

	//Given that the place in memory will always be different, we have to verify type rather than data
	suite.userServiceMock.AssertCalled(suite.T(), "CreateUser", mock.AnythingOfType(fmt.Sprintf("%T", &mockUser)))
	suite.userServiceMock.AssertNumberOfCalls(suite.T(), "CreateUser", 1)
}

// When failing to create a user, return an internal server error
func (suite *UsersControllerUnitTestSuite) TestCreateUser_RetrurnsInternalServerError() {

	test_utils.SetRequestBody(models.User{
		Email:    "some email",
		Password: "some password",
	}, suite.mockContext)

	suite.userServiceMock.On("CreateUser", mock.Anything).Return(errors.New("test"))

	suite.controller.CreateUser(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusInternalServerError, response.StatusCode)
}

func (suite *UsersControllerUnitTestSuite) TestCreateUser_RetrurnsCreated() {

	test_utils.SetRequestBody(models.User{
		Email:    "some email",
		Password: "some password",
	}, suite.mockContext)

	suite.userServiceMock.On("CreateUser", mock.Anything).Return(nil)

	suite.controller.CreateUser(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusCreated, response.StatusCode)
}

// When provided an invalid payload, should return bad request
func (suite *UsersControllerUnitTestSuite) TestLogin_ReturnsBadRequest() {

	test_utils.SetRequestBody(models.User{
		Email: "some email",
	}, suite.mockContext)

	suite.controller.Login(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusBadRequest, response.StatusCode)
}

func (suite *UsersControllerUnitTestSuite) TestLogin_AttemptsToValidateCredentials() {

	var expectedUser = models.User{
		Email:    "some email",
		Password: "some password",
	}
	test_utils.SetRequestBody(expectedUser, suite.mockContext)

	suite.userServiceMock.On("ValidateCredentials", mock.Anything).Return(false, errors.New("test"))

	suite.controller.Login(suite.mockContext)

	//Given that the place in memory will always be different, we have to verify type rather than data
	suite.userServiceMock.AssertCalled(suite.T(), "ValidateCredentials", mock.AnythingOfType(fmt.Sprintf("%T", &expectedUser)))
	suite.userServiceMock.AssertNumberOfCalls(suite.T(), "ValidateCredentials", 1)
}

// When failing to authenticate for a reason other than invalid credentials, return an internal server error
func (suite *UsersControllerUnitTestSuite) TestLogin_RetrurnsInternalServerError() {

	test_utils.SetRequestBody(models.User{
		Email:    "some email",
		Password: "some password",
	}, suite.mockContext)

	suite.userServiceMock.On("ValidateCredentials", mock.Anything).Return(false, errors.New("test"))

	suite.controller.Login(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusInternalServerError, response.StatusCode)
}

// When failing to authenticated due to invalid credentials, returns unauthorized
func (suite *UsersControllerUnitTestSuite) TestLogin_ReturnsUnauthorized() {

	test_utils.SetRequestBody(models.User{
		Email:    "some email",
		Password: "some password",
	}, suite.mockContext)

	suite.userServiceMock.On("ValidateCredentials", mock.Anything).Return(false, nil)

	suite.controller.Login(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *UsersControllerUnitTestSuite) TestLogin_AttemptsToCreateAnAuthToken() {

	var mockUser = models.User{
		//Note: Id is 0/ default number, since serialization ignores this, and will always be 0
		Id:       0,
		Email:    "some email",
		Password: "some password",
	}

	test_utils.SetRequestBody(mockUser, suite.mockContext)

	suite.userServiceMock.On("ValidateCredentials", mock.Anything).Return(true, nil)
	suite.jwtAuthorizerMock.On("GenerateToken", mock.Anything, mock.Anything).Return("", errors.New("test"))

	suite.controller.Login(suite.mockContext)

	suite.jwtAuthorizerMock.AssertCalled(suite.T(), "GenerateToken", mockUser.Email, strconv.FormatInt(mockUser.Id, 10))
	suite.jwtAuthorizerMock.AssertNumberOfCalls(suite.T(), "GenerateToken", 1)
}

// When failing to generat an auth token, return internal server error
func (suite *UsersControllerUnitTestSuite) TestLogin_ReturnsInternalServerError() {

	test_utils.SetRequestBody(models.User{
		Email:    "some email",
		Password: "some password",
	}, suite.mockContext)

	suite.userServiceMock.On("ValidateCredentials", mock.Anything).Return(true, nil)
	suite.jwtAuthorizerMock.On("GenerateToken", mock.Anything, mock.Anything).Return("", errors.New("test"))

	suite.controller.Login(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusInternalServerError, response.StatusCode)
}

func (suite *UsersControllerUnitTestSuite) TestLogin_ReturnsOk() {

	test_utils.SetRequestBody(models.User{
		Email:    "some email",
		Password: "some password",
	}, suite.mockContext)

	suite.userServiceMock.On("ValidateCredentials", mock.Anything).Return(true, nil)

	var expectedAuthToken string = "auth token"
	suite.jwtAuthorizerMock.On("GenerateToken", mock.Anything, mock.Anything).Return(expectedAuthToken, nil)

	suite.controller.Login(suite.mockContext)

	response := test_utils.GetHttpResponse(suite.mockResponseWriter)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Contains(response.Body, expectedAuthToken)
}
