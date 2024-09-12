package interfaces

import "example.com/models"

type IEventService interface {
	SaveEvent(event *models.Event) error
	GetEvents() ([]models.Event, error)
	GetEventById(id int64) (*models.Event, error)
	UpdateEvent(id int64, event models.Event) error
	DeleteEvent(id int64) error
}
