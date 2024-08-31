package person

import (
	"errors"
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

func New(firstName, lastName, birthDate string) (*Person, error) {

	if firstName == "" ||
		lastName == "" ||
		birthDate == "" {
		return nil, errors.New("First name, last name, and birth date are required")
	}

	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil
}

func (p Person) ShowUserData() {
	fmt.Printf("First Name: %s\n Last Name: %s\n Birth Date: %s\n Created At: %s\n",
		p.FirstName,
		p.LastName,
		p.BirthDate,
		p.CreatedAt.Format(time.DateTime))
}

func (p *Person) ClearUserName() {
	p.FirstName = ""
	p.LastName = ""
}
