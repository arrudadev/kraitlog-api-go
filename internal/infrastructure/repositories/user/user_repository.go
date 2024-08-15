package repository

import (
	"database/sql"

	"github.com/arrudadev/kraitlog-api/internal/domain/user/entity"
	"github.com/arrudadev/kraitlog-api/internal/domain/user/repository"
)

type UserRepositoryImplementation struct {
	db *sql.DB
}

func NewUserRepositoryImplementation(db *sql.DB) repository.UserRepository {
	return &UserRepositoryImplementation{db: db}
}

func (repository *UserRepositoryImplementation) Create(user *entity.User) error {
	_, err := repository.db.Exec("INSERT INTO USER(name, email, password) VALUES($1, $2, $3)", user.Name, user.Email, user.Password)

	return err
}

func (repository *UserRepositoryImplementation) FindByEmail(email string) (*entity.User, error) {
	row := repository.db.QueryRow("SELECT id, name, email FROM USER WHERE email = $1", email)

	user := entity.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
