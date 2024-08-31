package structs

import (
	"example.com/playground/lib"
	"example.com/playground/structs/person"
)

func GetUserData() (*person.Person, error) {
	var firstName, lastName, birthDate string

	lib.GetUserInput("Please enter your first name: ", &firstName)
	lib.GetUserInput("Please enter your last name: ", &lastName)
	lib.GetUserInput("Please enter your birthdate (MM/DD/YYYY): ", &birthDate)

	person, err := person.New(firstName, lastName, birthDate)

	if err != nil {
		return nil, err
	}

	return person, nil
}
