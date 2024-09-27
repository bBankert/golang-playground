package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"example.com/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type EventRepositoryUnitTestSuite struct {
	suite.Suite
	//Database mock "connection", do not use for interacting with the db, use "dbMock"
	database *sql.DB
	//Mock of the database that should be used to assert and interact with the database
	dbMock     sqlmock.Sqlmock
	repository *EventRepository
}

func TestEventRepositoryUnitTestSuite(t *testing.T) {
	suite.Run(t, &EventRepositoryUnitTestSuite{})
}

func (suite *EventRepositoryUnitTestSuite) SetupTest() {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		panic(fmt.Sprintf("Unable to create database, tests cannot proceed, error: %v\n", err.Error()))
	}

	suite.database = db

	suite.dbMock = mock

	suite.repository = NewEventRepository(db)
}

func (suite *EventRepositoryUnitTestSuite) TearDownTest() {

	//manually closing db connection, since using defer will close the connection
	//prior to starting the test
	suite.database.Close()
}

func (suite *EventRepositoryUnitTestSuite) TestAddEvent_PreparesTheSqlStatement() {

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")

	expectedEvent := models.Event{
		Id:          123,
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
		UserId:      1,
	}

	suite.dbMock.ExpectPrepare(`INSERT INTO Events (
	name,
	description,
	location,
	date,
	user_id
	) VALUES (?,?,?,?,?)`).
		ExpectExec().
		WithArgs(
			expectedEvent.Name,
			expectedEvent.Description,
			expectedEvent.Location,
			expectedEvent.Date,
			expectedEvent.UserId,
		).
		WillReturnResult(sqlmock.NewResult(int64(10), int64(1)))

	suite.repository.AddEvent(&expectedEvent)
}

// When an error occurs when preparing / executing the sql, will return the error
func (suite *EventRepositoryUnitTestSuite) TestAddEvent_ReturnsError() {

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")

	expectedEvent := models.Event{
		Id:          123,
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
		UserId:      1,
	}

	expectedError := errors.New("test")

	suite.dbMock.ExpectPrepare(`INSERT INTO Events (
	name,
	description,
	location,
	date,
	user_id
	) VALUES (?,?,?,?,?)`).
		ExpectExec().
		WithArgs(
			expectedEvent.Name,
			expectedEvent.Description,
			expectedEvent.Location,
			expectedEvent.Date,
			expectedEvent.UserId,
		).WillReturnError(expectedError)

	err := suite.repository.AddEvent(&expectedEvent)

	suite.NotNil(err)
	suite.Equal(expectedError, err)

}

func (suite *EventRepositoryUnitTestSuite) TestAddEvent_SetsTheIdToTheDbId() {

	var expectedId int64 = 10

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")

	expectedEvent := models.Event{
		Id:          45,
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
		UserId:      1,
	}

	suite.dbMock.ExpectPrepare(`INSERT INTO Events (
	name,
	description,
	location,
	date,
	user_id
	) VALUES (?,?,?,?,?)`).
		ExpectExec().
		WithArgs(
			expectedEvent.Name,
			expectedEvent.Description,
			expectedEvent.Location,
			expectedEvent.Date,
			expectedEvent.UserId,
		).
		WillReturnResult(sqlmock.NewResult(expectedId, int64(1)))

	suite.repository.AddEvent(&expectedEvent)

	suite.Equal(expectedId, expectedEvent.Id)

}

func (suite *EventRepositoryUnitTestSuite) TestAddEvent_ReturnsNil() {

	var expectedId int64 = 10

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")

	expectedEvent := models.Event{
		Id:          45,
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
		UserId:      1,
	}

	suite.dbMock.ExpectPrepare(`INSERT INTO Events (
	name,
	description,
	location,
	date,
	user_id
	) VALUES (?,?,?,?,?)`).
		ExpectExec().
		WithArgs(
			expectedEvent.Name,
			expectedEvent.Description,
			expectedEvent.Location,
			expectedEvent.Date,
			expectedEvent.UserId,
		).
		WillReturnResult(sqlmock.NewResult(expectedId, int64(1)))

	result := suite.repository.AddEvent(&expectedEvent)

	suite.Nil(result)

}

func (suite *EventRepositoryUnitTestSuite) TestGetEvents_PreparesTheSqlStatement() {

	suite.dbMock.ExpectQuery("SELECT * FROM Events").
		WillReturnRows(sqlmock.NewRows(make([]string, 0)))

	suite.repository.GetEvents()
}

// When an error occurs when preparing / executing the sql, will return the error
func (suite *EventRepositoryUnitTestSuite) TestGetEvents_ReturnsError() {

	expectedError := errors.New("test")

	suite.dbMock.ExpectQuery("SELECT * FROM Events").
		WillReturnError(expectedError)

	_, err := suite.repository.GetEvents()

	suite.NotNil(err)
	suite.Equal(expectedError, err)

}

// When no events exist, default to an empty array
func (suite *EventRepositoryUnitTestSuite) TestGetEvents_ReturnsEmptyArray() {

	suite.dbMock.ExpectQuery("SELECT * FROM Events").
		WillReturnRows(sqlmock.NewRows(make([]string, 0)))

	rows, _ := suite.repository.GetEvents()

	suite.Equal(len(rows), 0)

}

func (suite *EventRepositoryUnitTestSuite) TestGetEvents_ReturnsEvents() {

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")

	expectedEvent := models.Event{
		Id:          123,
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
		UserId:      1,
	}

	mockResult := sqlmock.NewRows([]string{
		"id",
		"name",
		"description",
		"location",
		"date",
		"user_id",
	}).AddRow(
		expectedEvent.Id,
		expectedEvent.Name,
		expectedEvent.Description,
		expectedEvent.Location,
		expectedEvent.Date,
		expectedEvent.UserId)

	suite.dbMock.ExpectQuery("SELECT * FROM Events").
		WillReturnRows(mockResult)

	rows, _ := suite.repository.GetEvents()

	suite.NotNil(rows)
	suite.Equal(1, len(rows))
	suite.Equal(expectedEvent, rows[0])

}

func (suite *EventRepositoryUnitTestSuite) TestGetEventById_PreparesTheSqlStatement() {

	var expectedId int64 = 123

	suite.dbMock.ExpectPrepare(`SELECT * FROM Events WHERE ID = ?`).
		ExpectQuery().
		WithArgs(expectedId).
		WillReturnRows(sqlmock.NewRows(make([]string, 0)))
	suite.repository.GetEventById(expectedId)
}

// When an error occurs when preparing / executing the sql, will return the error
func (suite *EventRepositoryUnitTestSuite) TestGetEventById_ReturnsError() {

	expectedError := errors.New("test")

	suite.dbMock.ExpectPrepare(`SELECT * FROM Events WHERE ID = ?`).
		ExpectQuery().
		WithArgs(int64(123)).
		WillReturnError(expectedError)

	_, err := suite.repository.GetEventById(int64(123))

	suite.NotNil(err)
	suite.Equal(expectedError, err)

}

func (suite *EventRepositoryUnitTestSuite) TestGetEventById_ReturnsTheEvent() {

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")

	expectedEvent := models.Event{
		Id:          123,
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
		UserId:      1,
	}

	mockResult := sqlmock.NewRows([]string{
		"id",
		"name",
		"description",
		"location",
		"date",
		"user_id",
	}).AddRow(
		expectedEvent.Id,
		expectedEvent.Name,
		expectedEvent.Description,
		expectedEvent.Location,
		expectedEvent.Date,
		expectedEvent.UserId)

	suite.dbMock.ExpectPrepare(`SELECT * FROM Events WHERE ID = ?`).
		ExpectQuery().
		WithArgs(int64(123)).
		WillReturnRows(mockResult)

	event, _ := suite.repository.GetEventById(int64(123))

	suite.NotNil(event)
	suite.Equal(&expectedEvent, event)

}

func (suite *EventRepositoryUnitTestSuite) TestUpdateEvent_PreparesTheSqlStatement() {

	var expectedId int64 = 123

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")

	expectedEvent := models.Event{
		Id:          123,
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
		UserId:      1,
	}

	suite.dbMock.ExpectPrepare(`UPDATE Events
	SET name = ?, description = ?, location = ?, date = ?, user_id = ?
	WHERE ID = ?`).
		ExpectExec().
		WithArgs(
			expectedEvent.Name,
			expectedEvent.Date,
			expectedEvent.Location,
			expectedEvent.Date,
			expectedEvent.UserId,
			expectedId).
		WillReturnResult(sqlmock.NewResult(int64(12), int64(1)))
	suite.repository.UpdateEvent(expectedId, expectedEvent)
}

// When an error occurs when preparing / executing the sql, will return the error
func (suite *EventRepositoryUnitTestSuite) TestUpdateEvent_ReturnsError() {

	expectedError := errors.New("test")

	var expectedId int64 = 123

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")

	expectedEvent := models.Event{
		Id:          123,
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
		UserId:      1,
	}

	suite.dbMock.ExpectPrepare(`UPDATE Events
	SET name = ?, description = ?, location = ?, date = ?, user_id = ?
	WHERE ID = ?`).
		ExpectExec().
		WithArgs(
			expectedEvent.Name,
			expectedEvent.Description,
			expectedEvent.Location,
			expectedEvent.Date,
			expectedEvent.UserId,
			expectedId).
		WillReturnError(expectedError)
	err := suite.repository.UpdateEvent(expectedId, expectedEvent)

	suite.NotNil(err)
	suite.Equal(expectedError, err)

}

func (suite *EventRepositoryUnitTestSuite) TestUpdateEvent_ReturnsNil() {

	var expectedId int64 = 123

	expectedDate, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00.000Z")

	expectedEvent := models.Event{
		Id:          123,
		Name:        "some name",
		Description: "some description",
		Location:    "some location",
		Date:        expectedDate,
		UserId:      1,
	}

	suite.dbMock.ExpectPrepare(`UPDATE Events
	SET name = ?, description = ?, location = ?, date = ?, user_id = ?
	WHERE ID = ?`).
		ExpectExec().
		WithArgs(
			expectedEvent.Name,
			expectedEvent.Description,
			expectedEvent.Location,
			expectedEvent.Date,
			expectedEvent.UserId,
			expectedId).
		WillReturnResult(sqlmock.NewResult(int64(123), int64(2)))
	err := suite.repository.UpdateEvent(expectedId, expectedEvent)

	suite.Nil(err)

}

func (suite *EventRepositoryUnitTestSuite) TestDeleteEvent_PreparesTheSqlStatement() {

	var expectedId int64 = 123

	suite.dbMock.ExpectPrepare(`DELETE FROM Events WHERE ID = ?`).
		ExpectExec().
		WithArgs(
			expectedId).
		WillReturnResult(sqlmock.NewResult(int64(12), int64(1)))
	suite.repository.DeleteEvent(expectedId)
}

// When an error occurs when preparing / executing the sql, will return the error
func (suite *EventRepositoryUnitTestSuite) TestDeleteEvent_ReturnsError() {

	expectedError := errors.New("test")

	var expectedId int64 = 123

	suite.dbMock.ExpectPrepare(`DELETE FROM Events WHERE ID = ?`).
		ExpectExec().
		WithArgs(
			expectedId).
		WillReturnError(expectedError)

	err := suite.repository.DeleteEvent(expectedId)

	suite.NotNil(err)
	suite.Equal(expectedError, err)
}

func (suite *EventRepositoryUnitTestSuite) TestDeleteEvent_ReturnsNil() {

	var expectedId int64 = 123

	suite.dbMock.ExpectPrepare(`DELETE FROM Events WHERE ID = ?`).
		ExpectExec().
		WithArgs(
			expectedId).
		WillReturnResult(sqlmock.NewResult(int64(12), int64(1)))

	err := suite.repository.DeleteEvent(expectedId)

	suite.Nil(err)

}
