package service

import (
	"errors"

	"github.com/arrudadev/kraitlog-api/internal/domain/user/entity"
	"github.com/arrudadev/kraitlog-api/internal/domain/user/repository"
)

type UserService interface {
	Create(firstName, lastName, email, password string) (*entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (service *userService) Create(firstName, lastName, email, password string) (*entity.User, error) {
	userExists, _ := service.userRepository.FindByEmail(email)
	if userExists != nil {
		return nil, errors.New("email already in use")
	}

	user, err := entity.NewUser(firstName, lastName, email, password)
	if err != nil {
		return nil, err
	}

	err = service.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
