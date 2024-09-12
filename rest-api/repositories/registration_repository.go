package repositories

import "database/sql"

type RegistrationRepository struct {
	database *sql.DB
}

func (registrationRepository RegistrationRepository) CreateRegistration(eventId, userId int64) error {
	createRegistrationSql := `
	INSERT INTO Registrations(event_id, user_id)
	VALUES (?, ?)`

	statement, err := registrationRepository.database.Prepare(createRegistrationSql)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, resultError := statement.Exec(eventId, userId)

	if resultError != nil {
		return resultError
	}

	return nil
}

func (registrationRepository RegistrationRepository) DeleteRegistration(eventId, userId int64) error {
	createRegistrationSql := `
	DELETE FROM Registrations
	WHERE event_id = ? AND user_id = ?`

	statement, err := registrationRepository.database.Prepare(createRegistrationSql)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, resultError := statement.Exec(eventId, userId)

	if resultError != nil {
		return resultError
	}

	return nil
}

func NewRegistrationRepository(database *sql.DB) *RegistrationRepository {
	return &RegistrationRepository{
		database: database,
	}
}
