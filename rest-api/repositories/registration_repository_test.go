package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type RegistrationRepositoryUnitTestSuite struct {
	suite.Suite
	//Database mock "connection", do not use for interacting with the db, use "dbMock"
	database *sql.DB
	//Mock of the database that should be used to assert and interact with the database
	dbMock     sqlmock.Sqlmock
	repository *RegistrationRepository
}

func TestRegistrationRepositoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &RegistrationRepositoryUnitTestSuite{})
}

func (suite *RegistrationRepositoryUnitTestSuite) SetupTest() {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		panic(fmt.Sprintf("Unable to create database, tests cannot proceed, error: %v\n", err.Error()))
	}

	suite.database = db

	suite.dbMock = mock

	suite.repository = NewRegistrationRepository(db)
}

func (suite *RegistrationRepositoryUnitTestSuite) TearDownTest() {

	//manually closing db connection, since using defer will close the connection
	//prior to starting the test
	suite.database.Close()
}

func (suite *RegistrationRepositoryUnitTestSuite) TestCreateRegistration_PreparesTheQuery() {

	var (
		expectedEventId int64 = 12
		expectedUserId  int64 = 13
	)

	suite.dbMock.ExpectPrepare(`
	INSERT INTO Registrations(event_id, user_id)
	VALUES (?, ?)`).
		ExpectExec().
		WithArgs(
			expectedEventId,
			expectedUserId,
		).
		WillReturnResult(sqlmock.NewResult(int64(10), int64(1)))

	suite.repository.CreateRegistration(expectedEventId, expectedUserId)
}

// When a db error occurs, pass that up to the caller
func (suite *RegistrationRepositoryUnitTestSuite) TestCreateRegistration_ReturnsTheError() {

	var (
		expectedEventId int64 = 12
		expectedUserId  int64 = 13
		expectedError   error = errors.New("test")
	)

	suite.dbMock.ExpectPrepare(`
	INSERT INTO Registrations(event_id, user_id)
	VALUES (?, ?)`).
		ExpectExec().
		WithArgs(
			expectedEventId,
			expectedUserId,
		).
		WillReturnError(expectedError)

	err := suite.repository.CreateRegistration(expectedEventId, expectedUserId)

	suite.NotNil(err)
	suite.Equal(expectedError, err)
}

func (suite *RegistrationRepositoryUnitTestSuite) TestCreateRegistration_ReturnsNil() {

	var (
		expectedEventId int64 = 12
		expectedUserId  int64 = 13
	)

	suite.dbMock.ExpectPrepare(`
	INSERT INTO Registrations(event_id, user_id)
	VALUES (?, ?)`).
		ExpectExec().
		WithArgs(
			expectedEventId,
			expectedUserId,
		).
		WillReturnResult(sqlmock.NewResult(int64(10), int64(1)))

	err := suite.repository.CreateRegistration(expectedEventId, expectedUserId)

	suite.Nil(err)
}

func (suite *RegistrationRepositoryUnitTestSuite) TestDeleteRegistration_PreparesTheQuery() {

	var (
		expectedEventId int64 = 12
		expectedUserId  int64 = 13
	)

	suite.dbMock.ExpectPrepare(`DELETE FROM Registrations
	WHERE event_id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			expectedEventId,
			expectedUserId,
		).
		WillReturnResult(sqlmock.NewResult(int64(10), int64(1)))

	suite.repository.DeleteRegistration(expectedEventId, expectedUserId)
}

// When a db error occurs, pass that up to the caller
func (suite *RegistrationRepositoryUnitTestSuite) TestDeleteRegistration_ReturnsTheError() {

	var (
		expectedEventId int64 = 12
		expectedUserId  int64 = 13
		expectedError   error = errors.New("test")
	)

	suite.dbMock.ExpectPrepare(`DELETE FROM Registrations
	WHERE event_id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			expectedEventId,
			expectedUserId,
		).
		WillReturnError(expectedError)

	err := suite.repository.DeleteRegistration(expectedEventId, expectedUserId)

	suite.NotNil(err)
	suite.Equal(expectedError, err)
}

func (suite *RegistrationRepositoryUnitTestSuite) TestDeleteRegistration_ReturnsNil() {

	var (
		expectedEventId int64 = 12
		expectedUserId  int64 = 13
	)

	suite.dbMock.ExpectPrepare(`DELETE FROM Registrations
	WHERE event_id = ? AND user_id = ?`).
		ExpectExec().
		WithArgs(
			expectedEventId,
			expectedUserId,
		).
		WillReturnResult(sqlmock.NewResult(int64(10), int64(1)))

	err := suite.repository.DeleteRegistration(expectedEventId, expectedUserId)

	suite.Nil(err)
}
