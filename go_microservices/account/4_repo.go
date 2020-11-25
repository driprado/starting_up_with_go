package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

//ErrRepo is the customised error for repo package
var ErrRepo = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepo returns a Repository interface which will be implemented in the repo struct above
func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{ // return a repo structure with the db and logger inside
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user User) error {
	sql := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)` // insert values that will be passed into table users

	if user.Email == "" || user.Password == "" {
		return ErrRepo // If email or password don't exist, return error
	}

	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password) // Execute sql line above
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email) // scan the result of the query bac to the variable email
	if err != nil {
		return "", ErrRepo
	}
	return email, nil
}
