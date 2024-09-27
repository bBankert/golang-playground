package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"example.com/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryUnitTestSuite struct {
	suite.Suite
	//Database mock "connection", do not use for interacting with the db, use "dbMock"
	database *sql.DB
	//Mock of the database that should be used to assert and interact with the database
	dbMock     sqlmock.Sqlmock
	repository *UserRepository
}

func TestUserRepositoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &UserRepositoryUnitTestSuite{})
}

func (suite *UserRepositoryUnitTestSuite) SetupTest() {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		panic(fmt.Sprintf("Unable to create database, tests cannot proceed, error: %v\n", err.Error()))
	}

	suite.database = db

	suite.dbMock = mock

	suite.repository = NewUserRepository(db)
}

func (suite *UserRepositoryUnitTestSuite) TearDownTest() {

	//manually closing db connection, since using defer will close the connection
	//prior to starting the test
	suite.database.Close()
}

func (suite *UserRepositoryUnitTestSuite) TestCreateUser_ShouldPrepareTheQuery() {

	expectedUser := models.User{
		Id:       123,
		Email:    "some email",
		Password: "some password",
	}

	suite.dbMock.ExpectPrepare(`
	INSERT INTO Users(email, password)
	VALUES (?, ?)`).
		ExpectExec().
		WithArgs(
			expectedUser.Email,
			expectedUser.Password,
		).
		WillReturnResult(sqlmock.NewResult(int64(10), int64(1)))

	suite.repository.CreateUser(&expectedUser)
}

// When an error occurs during db interaction, pass it up
func (suite *UserRepositoryUnitTestSuite) TestCreateUser_ReturnsTheError() {

	expectedError := errors.New("test")

	expectedUser := models.User{
		Id:       123,
		Email:    "some email",
		Password: "some password",
	}

	suite.dbMock.ExpectPrepare(`
	INSERT INTO Users(email, password)
	VALUES (?, ?)`).
		ExpectExec().
		WithArgs(
			expectedUser.Email,
			expectedUser.Password,
		).
		WillReturnError(expectedError)

	err := suite.repository.CreateUser(&expectedUser)

	suite.NotNil(err)
	suite.Equal(expectedError, err)
}

func (suite *UserRepositoryUnitTestSuite) TestCreateUser_ShouldUpdateTheUserWithTheDbId() {

	expectedUser := models.User{
		Id:       123,
		Email:    "some email",
		Password: "some password",
	}

	expectedId := int64(50)

	suite.dbMock.ExpectPrepare(`
	INSERT INTO Users(email, password)
	VALUES (?, ?)`).
		ExpectExec().
		WithArgs(
			expectedUser.Email,
			expectedUser.Password,
		).
		WillReturnResult(sqlmock.NewResult(expectedId, int64(1)))

	suite.repository.CreateUser(&expectedUser)

	suite.Equal(expectedId, expectedUser.Id)
}

func (suite *UserRepositoryUnitTestSuite) TestCreateUser_ReturnsNil() {

	expectedUser := models.User{
		Id:       123,
		Email:    "some email",
		Password: "some password",
	}

	suite.dbMock.ExpectPrepare(`
	INSERT INTO Users(email, password)
	VALUES (?, ?)`).
		ExpectExec().
		WithArgs(
			expectedUser.Email,
			expectedUser.Password,
		).
		WillReturnResult(sqlmock.NewResult(int64(50), int64(1)))

	err := suite.repository.CreateUser(&expectedUser)

	suite.Nil(err)
}

func (suite *UserRepositoryUnitTestSuite) TestGetUserByEmail_ShouldPrepareTheQuery() {

	expectedEmail := "some email"

	suite.dbMock.ExpectPrepare(`SELECT * FROM Users WHERE email = ?`).
		ExpectQuery().
		WithArgs(
			expectedEmail,
		).
		WillReturnRows(sqlmock.NewRows(make([]string, 0)))

	suite.repository.GetUserByEmail(expectedEmail)
}

// When an error occurs during db interaction, pass it up
func (suite *UserRepositoryUnitTestSuite) TestGetUserByEmail_ReturnsTheError() {

	expectedEmail := "some email"
	expectedError := errors.New("test")

	suite.dbMock.ExpectPrepare(`SELECT * FROM Users WHERE email = ?`).
		ExpectQuery().
		WithArgs(
			expectedEmail,
		).
		WillReturnError(expectedError)

	_, err := suite.repository.GetUserByEmail(expectedEmail)

	suite.NotNil(err)
	suite.Equal(expectedError, err)
}

func (suite *UserRepositoryUnitTestSuite) TestGetUserByEmail_ReturnsTheUser() {

	expectedUser := models.User{
		Id:       123,
		Email:    "some email",
		Password: "some password",
	}

	mockResult := sqlmock.NewRows([]string{
		"id",
		"email",
		"password",
	}).AddRow(
		expectedUser.Id,
		expectedUser.Email,
		expectedUser.Password,
	)

	suite.dbMock.ExpectPrepare(`SELECT * FROM Users WHERE email = ?`).
		ExpectQuery().
		WithArgs(
			expectedUser.Email,
		).
		WillReturnRows(mockResult)

	user, _ := suite.repository.GetUserByEmail(expectedUser.Email)

	suite.Equal(&expectedUser, user)
}
