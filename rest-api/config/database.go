package config

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func InitializeDatabase() *sql.DB {
	var err error
	var database *sql.DB
	database, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Unable to connect to the database")
	}

	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(5)

	setupTables(database)

	return database
}

func CloseDatabaseConnection(database *sql.DB) {
	err := database.Close()

	if err != nil {
		panic(err)
	}
}

func setupTables(database *sql.DB) {

	createUserTableSql := `
	CREATE TABLE IF NOT EXISTS Users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err := database.Exec(createUserTableSql)

	if err != nil {
		panic("Unable to create user table")
	}

	createEventsTableSql := `
	CREATE TABLE IF NOT EXISTS Events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES Users(id)
	)`

	_, err = database.Exec(createEventsTableSql)

	if err != nil {
		panic("Unable to create table")
	}

	createRegistrationsTableSql := `
	CREATE TABLE IF NOT EXISTS Registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES Events(id)
		FOREIGN KEY(user_id) REFERENCES User(id)
	)`

	_, err = database.Exec(createRegistrationsTableSql)

	if err != nil {
		panic("Unable to create registrations table")
	}
}
