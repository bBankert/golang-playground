package services

import (
	libInterfaces "example.com/interfaces/lib"
	repositoryInterfaces "example.com/interfaces/repositories"
	"example.com/models"
)

type UserService struct {
	userRepository repositoryInterfaces.IUserRepository
	passwordHasher libInterfaces.IHasher
}

func (userService UserService) CreateUser(user *models.User) error {

	hashedPassword, err := userService.passwordHasher.HashPassword(user.Password)

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

	return userService.passwordHasher.ValidatePasswordHash(
		user.Password,
		savedUser.Password), nil
}

func NewUserService(userRepository repositoryInterfaces.IUserRepository, passwordHasher libInterfaces.IHasher) *UserService {
	return &UserService{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
	}
}
