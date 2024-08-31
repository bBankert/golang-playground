package admin

import (
	"time"

	"example.com/playground/structs/person"
)

type Admin struct {
	email    string
	password string
	Person   person.Person
}

func New(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		Person: person.Person{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			BirthDate: "---",
			CreatedAt: time.Now(),
		},
	}
}
