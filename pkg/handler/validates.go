package handler

import (
	"Ozinshe/pkg/repository"
	"errors"
	"regexp"

	"github.com/jmoiron/sqlx"
)

func validateUserData(user repository.SignupUser, db *sqlx.DB) error {
	if !isValidEmail(user.Email) {
		return errors.New("некорректный формат email")
	}
	checkEmail := checkEmailExists(user.Email)
	exists, err := checkEmail(db)
	if err != nil {
		return errors.New("ошибка при проверке email")
	}
	if exists {
		return errors.New("пользователь с таким email уже существует")
	}
	isValid, err := isValidPassword(user.Password, user.ConfirmPassword)
	if !isValid {
		return errors.New(err.Error())
	}
	return nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidPassword(password, confirm_password string) (bool, error) {
	if len(password) < 8 {
		return false, errors.New("пароль должен быть не менее 8 символов")
	}
	if password != confirm_password {
		return false, errors.New("пароли не совпадают")
	}
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	if !uppercaseRegex.MatchString(password) {
		return false, errors.New("пароль должен содержать хотя бы одну заглавную букву")
	}
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	if !lowercaseRegex.MatchString(password) {
		return false, errors.New("пароль должен содержать хотя бы одну строчную букву")
	}
	digitRegex := regexp.MustCompile(`\d`)
	if !digitRegex.MatchString(password) {
		return false, errors.New("пароль должен содержать хотя бы одну цифру")
	}
	return true, errors.New("")
}

var ErrNoRows = errors.New("no rows found")

func checkEmailExists(email string) func(*sqlx.DB) (bool, error) {
	return func(db *sqlx.DB) (bool, error) {
		var count int
		query := "Select COUNT(*) from users where email = $1"
		row := db.QueryRow(query, email)
		err := row.Scan(&count)
		if err != nil {
			return false, err
		}
		return count > 0, nil
	}
}
