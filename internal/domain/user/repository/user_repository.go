package repository

import "github.com/arrudadev/kraitlog-api/internal/domain/user/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
