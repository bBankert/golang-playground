package interfaces

type IRegistrationRepository interface {
	CreateRegistration(eventId, userId int64) error
	DeleteRegistration(eventId, userId int64) error
}
