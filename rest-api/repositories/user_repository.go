package repositories

import (
	"database/sql"
	"errors"

	"example.com/models"
)

type UserRepository struct {
	database *sql.DB
}

func (userRepository UserRepository) CreateUser(user *models.User) error {
	createUserSql := `
	INSERT INTO Users(email, password)
	VALUES (?, ?)`

	statement, err := userRepository.database.Prepare(createUserSql)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, creationError := statement.Exec(user.Email, user.Password)

	if creationError != nil {
		return creationError
	}

	id, _ := result.LastInsertId()

	user.Id = id

	return nil
}

func (userRepository UserRepository) GetUserByEmail(email string) (*models.User, error) {
	selectUserByEmailSql := `SELECT * FROM Users WHERE email = ?`

	statement, err := userRepository.database.Prepare(selectUserByEmailSql)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, queryError := statement.Query(email)

	if queryError != nil {
		return nil, queryError
	}

	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("no user found associated with provided email address")
	}

	var user models.User

	err = rows.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{
		database: database,
	}
}
