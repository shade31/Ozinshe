package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
	CreateUser(user SignupUser) (int, error)
	GetUser(email, password string) (User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
