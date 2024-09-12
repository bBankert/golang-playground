package repositories

import (
	"database/sql"

	"example.com/models"
)

type EventRepository struct {
	database *sql.DB
}

func (eventRepository *EventRepository) AddEvent(event *models.Event) error {
	saveSql := `
	INSERT INTO Events (
	name,
	description,
	location,
	date,
	user_id
	) VALUES (?,?,?,?,?)`

	statement, err := eventRepository.database.Prepare(saveSql)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, resultError := statement.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.Date,
		event.UserId)

	if resultError != nil {
		return resultError
	}

	id, _ := result.LastInsertId()

	event.Id = id

	return nil
}

func (eventRepository *EventRepository) GetEvents() ([]models.Event, error) {
	eventsQuerySql := "SELECT * FROM Events"

	rows, err := eventRepository.database.Query(eventsQuerySql)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var event models.Event

		err = rows.Scan(
			&event.Id,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.Date,
			&event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	//preventing null responses to default to empty arrays for better error handling
	if events == nil {
		events = make([]models.Event, 0)
	}

	return events, nil
}

func (eventRepository *EventRepository) GetEventById(id int64) (*models.Event, error) {
	eventByIdQuerySql := "SELECT * FROM Events WHERE ID = ?"

	statement, err := eventRepository.database.Prepare(eventByIdQuerySql)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, queryResult := statement.Query(id)

	if queryResult != nil {
		return nil, queryResult
	}

	defer rows.Close()

	var event models.Event

	if rows.Next() {
		err = rows.Scan(
			&event.Id,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.Date,
			&event.UserId)

		if err != nil {
			return nil, err
		}
	}

	return &event, nil
}

func (eventRepository *EventRepository) UpdateEvent(id int64, event models.Event) error {
	updateEventSql := `
	UPDATE Events
	SET name = ?, description = ?, location = ?, date = ?, user_id = ?
	WHERE ID = ?`

	statement, err := eventRepository.database.Prepare(updateEventSql)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, updateError := statement.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.Date,
		event.UserId,
		id)

	if updateError != nil {
		return updateError
	}

	return nil
}

func (eventRepository *EventRepository) DeleteEvent(id int64) error {
	deleteEventSql := `DELETE FROM Events WHERE ID = ?`

	statement, err := eventRepository.database.Prepare(deleteEventSql)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, deleteError := statement.Exec(id)

	if deleteError != nil {
		return deleteError
	}

	return nil
}

func NewEventRepository(database *sql.DB) *EventRepository {
	return &EventRepository{
		database: database,
	}
}
