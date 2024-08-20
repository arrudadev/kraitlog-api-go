package handler

import (
	"net/http"

	"github.com/arrudadev/kraitlog-api/internal/application/user/dto"
	usecase "github.com/arrudadev/kraitlog-api/internal/application/user/usecase/create_user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUserUseCase *usecase.CreateUserUseCase
}

func NewUserHandler(createUserUseCase *usecase.CreateUserUseCase) *UserHandler {
	return &UserHandler{createUserUseCase: createUserUseCase}
}

func (handler *UserHandler) CreateUser(context *gin.Context) {
	var createUserDTO *dto.CreateUserDTO

	if err := context.ShouldBindJSON(&createUserDTO); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDTO, err := handler.createUserUseCase.Execute(createUserDTO)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": userDTO})
}
