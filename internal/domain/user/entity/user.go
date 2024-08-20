package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func NewUser(firstName, lastName, email, password string) (*User, error) {
	if firstName == "" || lastName == "" || email == "" || password == "" {
		return nil, errors.New("all user fields are required")
	}

	hashedPassword, err := hashPassword(password)
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

func (user *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) FullName() string {
	return user.FirstName + " " + user.LastName
}
