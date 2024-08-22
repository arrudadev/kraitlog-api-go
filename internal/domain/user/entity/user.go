package entity

import (
	"errors"
	"time"

	"github.com/arrudadev/kraitlog-api/internal/domain/auth/service"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(firstName, lastName, email, password string) (*User, error) {
	if firstName == "" || lastName == "" || email == "" || password == "" {
		return nil, errors.New("all user fields are required")
	}

	authService := service.NewAuthService()
	hashedPassword, err := authService.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}, nil
}

func (user *User) FullName() string {
	return user.FirstName + " " + user.LastName
}
