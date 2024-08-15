package usecase

import (
	"github.com/arrudadev/kraitlog-api/internal/application/user/dto"
	"github.com/arrudadev/kraitlog-api/internal/domain/user/entity"
	"github.com/arrudadev/kraitlog-api/internal/domain/user/service"
)

type CreateUserUseCase struct {
	userService service.UserService
}

func NewCreateUserUseCase(userService service.UserService) *CreateUserUseCase {
	return &CreateUserUseCase{
		userService: userService,
	}
}

func (useCase *CreateUserUseCase) Execute(dto *dto.CreateUserDTO) (*entity.User, error) {
	return useCase.userService.Create(dto.Name, dto.Email, dto.Password)
}
