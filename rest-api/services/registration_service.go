package services

import (
	"errors"

	"example.com/constants"
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
		return errors.New(constants.NO_EVENT_FOR_ID_ERROR)
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
		return errors.New(constants.NO_EVENT_FOR_ID_ERROR)
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
