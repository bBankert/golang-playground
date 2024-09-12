package interfaces

type IRegistrationService interface {
	CreateRegistration(eventId, userId int64) error
	DeleteRegistration(eventId, userId int64) error
}
