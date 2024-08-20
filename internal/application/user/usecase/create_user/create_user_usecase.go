package usecase

import (
	"github.com/arrudadev/kraitlog-api/internal/application/user/dto"
	"github.com/arrudadev/kraitlog-api/internal/domain/user/service"
	"github.com/arrudadev/kraitlog-api/internal/shared/utils"
)

type CreateUserUseCase struct {
	userService service.UserService
}

func NewCreateUserUseCase(userService service.UserService) *CreateUserUseCase {
	return &CreateUserUseCase{
		userService: userService,
	}
}

func (useCase *CreateUserUseCase) Execute(inputDTO *dto.CreateUserDTO) (*dto.UserDTO, error) {
	user, err := useCase.
		userService.Create(
		inputDTO.FirstName,
		inputDTO.LastName,
		inputDTO.Email,
		inputDTO.Password)
	if err != nil {
		return nil, err
	}

	userDTO := &dto.UserDTO{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		FullName:  user.FullName(),
		Email:     user.Email,
		CreatedAt: utils.FormatDateTimeUTC(user.CreatedAt),
	}

	return userDTO, nil
}
