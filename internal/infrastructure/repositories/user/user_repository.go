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
	_, err := repository.
		db.Exec(`
			INSERT INTO users(id, first_name, last_name, email, password, created_at) 
			VALUES($1, $2, $3, $4, $5, $6)
		`,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt)

	return err
}

func (repository *UserRepositoryImplementation) FindByEmail(email string) (*entity.User, error) {
	row := repository.
		db.QueryRow(`
			SELECT 
				id, 
				first_name, 
				last_name, 
				email, 
				created_at, 
				updated_at 
			FROM users WHERE email = $1
		`,
		email)

	user := entity.User{}
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
