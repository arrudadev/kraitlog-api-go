package routes

import (
	"database/sql"

	createUserUseCase "github.com/arrudadev/kraitlog-api/internal/application/user/usecase/create_user"
	"github.com/arrudadev/kraitlog-api/internal/domain/user/service"
	api "github.com/arrudadev/kraitlog-api/internal/infrastructure/api/user/handler"
	repository "github.com/arrudadev/kraitlog-api/internal/infrastructure/repositories/user"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, db *sql.DB) {
	userRepository := repository.NewUserRepositoryImplementation(db)
	userService := service.NewUserService(userRepository)
	createUserUseCase := createUserUseCase.NewCreateUserUseCase(userService)
	userHandler := api.NewUserHandler(createUserUseCase)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", userHandler.CreateUser)
	}
}
