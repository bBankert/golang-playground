package services

import (
	"errors"

	interfaces "example.com/interfaces/repositories"
)

type RegistrationService struct {
	registrationRepository interfaces.IRegistrationRepository
	eventRepository        interfaces.IEventRepository
}

func (registrationService RegistrationService) CreateRegistration(eventId, userId int64) error {
	event, err := registrationService.eventRepository.GetEventById(eventId)

	if err != nil {
		return err
	} else if event.Id == 0 {
		return errors.New("no event exists with provided id")
	}

	err = registrationService.registrationRepository.CreateRegistration(eventId, userId)

	if err != nil {
		return err
	}

	return nil
}

func (registrationService RegistrationService) DeleteRegistration(eventId, userId int64) error {

	event, err := registrationService.eventRepository.GetEventById(eventId)

	if err != nil {
		return err
	} else if event.Id == 0 {
		return errors.New("no event exists with provided id")
	}

	err = registrationService.registrationRepository.DeleteRegistration(eventId, userId)

	if err != nil {
		return err
	}

	return nil

}

func NewRegistrationService(
	registrationRepository interfaces.IRegistrationRepository,
	eventRepository interfaces.IEventRepository) *RegistrationService {
	return &RegistrationService{
		registrationRepository: registrationRepository,
		eventRepository:        eventRepository,
	}
}
