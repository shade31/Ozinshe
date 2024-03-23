package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user SignupUser) (int, error) {
	var id int
	query := fmt.Sprint("insert into users(email, password) values($1, $2) returning id")

	row := r.db.QueryRow(query, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// func (r *AuthPostgres) GetUser(username, password string) (User, error) {
// 	var user User
// 	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
// 	err := r.db.Get(&user, query, username, password)

// 	return user, err
// }

// func (r *AuthPostgres) GetUser(username, password string) (User, error) {

// }
