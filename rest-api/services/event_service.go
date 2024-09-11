package services

import (
	interfaces "example.com/interfaces/repositories"
	"example.com/models"
)

type EventService struct {
	eventRepository interfaces.IEventRepository
}

func (eventService EventService) SaveEvent(event *models.Event) error {
	err := eventService.eventRepository.AddEvent(event)

	if err != nil {
		return err
	}

	return nil
}

func (eventService EventService) GetEvents() ([]models.Event, error) {
	events, err := eventService.eventRepository.GetEvents()

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (eventService EventService) GetEventById(id int64) (*models.Event, error) {
	event, err := eventService.eventRepository.GetEventById(id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (eventService EventService) UpdateEvent(id int64, event models.Event) error {
	err := eventService.eventRepository.UpdateEvent(id, event)

	if err != nil {
		return err
	}

	return nil
}

func (eventService EventService) DeleteEvent(id int64) error {
	err := eventService.eventRepository.DeleteEvent(id)

	if err != nil {
		return err
	}

	return nil
}

func NewEventService(eventRepository interfaces.IEventRepository) *EventService {
	return &EventService{
		eventRepository: eventRepository,
	}
}
