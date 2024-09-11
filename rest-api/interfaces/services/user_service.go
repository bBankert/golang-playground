package interfaces

import "example.com/models"

type IUserService interface {
	CreateUser(user *models.User) error
	ValidateCredentials(user *models.User) (bool, error)
}
