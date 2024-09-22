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

type UserServiceUnitTestSuite struct {
	suite.Suite
	passwordHasherMock mocks.IHasher
	userRepositoryMock mocks.IUserRepository
	service            *UserService
}

func TestUserServiceUnitTestSuite(t *testing.T) {
	suite.Run(t, &EventServiceUnitTestSuite{})
}

func (suite *UserServiceUnitTestSuite) SetupTest() {
	suite.userRepositoryMock = mocks.IUserRepository{}
	suite.passwordHasherMock = mocks.IHasher{}

	suite.service = NewUserService(
		&suite.userRepositoryMock,
		&suite.passwordHasherMock)
}

func (suite *UserServiceUnitTestSuite) TestCreateUser_AttemptsToHashTheRawPassword() {

	expectedPassword := "some password"

	suite.userRepositoryMock.On("HashPassword", mock.Anything).Return(nil, errors.New("test"))

	suite.service.CreateUser(&models.User{
		Password: expectedPassword,
	})

	suite.userRepositoryMock.AssertCalled(suite.T(), "HashPassword", expectedPassword)
	suite.userRepositoryMock.AssertNumberOfCalls(suite.T(), "HashPassword", 1)
}

// When failing to hash the password, pass up the error
func (suite *UserServiceUnitTestSuite) TestCreateUser_ReturnsAnError() {

	expectedError := errors.New("test")

	suite.userRepositoryMock.On("HashPassword", mock.Anything).Return(nil, expectedError)

	err := suite.service.CreateUser(&models.User{
		Password: "some password",
	})

	suite.NotNil(err)
	suite.Equal(err, expectedError)
}

func (suite *UserServiceUnitTestSuite) TestCreateUser_AttemptsToCreateTheUser() {

	suite.userRepositoryMock.On("HashPassword", mock.Anything).Return("some hashed password", nil)
	suite.userRepositoryMock.On("CreateUser", mock.Anything).Return(errors.New("test"))

	suite.service.CreateUser(&models.User{
		Password: "some password",
	})

	suite.userRepositoryMock.AssertCalled(suite.T(), "CreateUser", mock.AnythingOfType(fmt.Sprintf("%T", &models.User{})))
	suite.userRepositoryMock.AssertNumberOfCalls(suite.T(), "CreateUser", 1)
}

// When db interaction fails, pass up the error
func (suite *UserServiceUnitTestSuite) TestCreateUserWhenFailingToCreateTheUser_ReturnsAnError() {

	expectedError := errors.New("test")

	suite.userRepositoryMock.On("HashPassword", mock.Anything).Return("some hashed password", nil)
	suite.userRepositoryMock.On("CreateUser", mock.Anything).Return(expectedError)

	err := suite.service.CreateUser(&models.User{
		Password: "some password",
	})

	suite.NotNil(err)
	suite.Equal(err, expectedError)
}

func (suite *UserServiceUnitTestSuite) TestCreateUser_ReturnsNil() {

	suite.userRepositoryMock.On("HashPassword", mock.Anything).Return("some hashed password", nil)
	suite.userRepositoryMock.On("CreateUser", mock.Anything).Return(nil)

	err := suite.service.CreateUser(&models.User{
		Password: "some password",
	})

	suite.Nil(err)
}

func (suite *UserServiceUnitTestSuite) TestValidateCredentials_AttemptsToFetchTheUserByEmail() {

	expectedUser := models.User{
		Password: "some password",
		Email:    "some email",
	}

	suite.userRepositoryMock.On("GetUserByEmail", mock.Anything).Return(nil, errors.New("test"))

	suite.service.ValidateCredentials(&expectedUser)

	suite.userRepositoryMock.AssertCalled(suite.T(), "GetUserByEmail", expectedUser.Email)
	suite.userRepositoryMock.AssertNumberOfCalls(suite.T(), "GetUserByEmail", 1)
}

// When an error occurs trying to fetch the user, pass it up
func (suite *UserServiceUnitTestSuite) TestValidateCredentials_ReturnsAnError() {

	expectedError := errors.New("test")

	suite.userRepositoryMock.On("GetUserByEmail", mock.Anything).Return(nil, expectedError)

	_, err := suite.service.ValidateCredentials(&models.User{
		Password: "some password",
		Email:    "some email",
	})

	suite.NotNil(err)
	suite.Equal(err, expectedError)
}

func (suite *UserServiceUnitTestSuite) TestValidateCredentialsWhenNoUserIsFound_ReturnsNil() {

	suite.userRepositoryMock.On("GetUserByEmail", mock.Anything).Return(nil, nil)

	validCredentials, err := suite.service.ValidateCredentials(&models.User{
		Password: "some password",
		Email:    "some email",
	})

	suite.False(validCredentials)
	suite.Nil(err)
}

func (suite *UserServiceUnitTestSuite) TestValidateCredentials_AttemptsToValidateTheCredentials() {

	rawPassword, savedPassword := "some password", "some saved password"

	suite.userRepositoryMock.On("GetUserByEmail", mock.Anything).Return(&models.User{
		Id:       1,
		Email:    "some email",
		Password: savedPassword,
	}, nil)

	suite.passwordHasherMock.On("ValidatePasswordHash", mock.Anything, mock.Anything).Return(false)

	suite.service.ValidateCredentials(&models.User{
		Password: rawPassword,
		Email:    "some email",
	})

	suite.passwordHasherMock.AssertCalled(suite.T(), "ValidatePasswordHash", rawPassword, savedPassword)
	suite.passwordHasherMock.AssertNumberOfCalls(suite.T(), "ValidatePasswordHash", 1)
}

func (suite *UserServiceUnitTestSuite) TestValidateCredentials_ReturnsTheValidationResult() {

	expectedValidationResult := true

	suite.userRepositoryMock.On("GetUserByEmail", mock.Anything).Return(&models.User{
		Id:       1,
		Email:    "some email",
		Password: "some password",
	}, nil)

	suite.passwordHasherMock.On("ValidatePasswordHash", mock.Anything, mock.Anything).Return(expectedValidationResult)

	validationResult, err := suite.service.ValidateCredentials(&models.User{
		Password: "some password",
		Email:    "some email",
	})

	suite.Nil(err)
	suite.Equal(expectedValidationResult, validationResult)
}
