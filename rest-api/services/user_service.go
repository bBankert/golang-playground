package services

import (
	interfaces "example.com/interfaces/repositories"
	"example.com/lib"
	"example.com/models"
)

type UserService struct {
	userRepository interfaces.IUserRepository
}

func (userService UserService) CreateUser(user *models.User) error {

	hashedPassword, err := lib.HashPassword(user.Password)

	if err != nil {
		return nil
	}

	user.Password = hashedPassword

	err = userService.userRepository.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (userService UserService) ValidateCredentials(user *models.User) (bool, error) {

	savedUser, err := userService.userRepository.GetUserByEmail(user.Email)

	if err != nil {
		return false, err
	}

	if savedUser.Id == 0 {
		return false, nil
	}

	user.Id = savedUser.Id

	return lib.ValidatePasswordHash(
		user.Password,
		savedUser.Password), nil
}

func NewUserService(userRepository interfaces.IUserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}
