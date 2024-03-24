package repository

type User struct {
	Id         int    `json:"-" db:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Birthdate  string `json:"birthdate"`
	Phone      string `json:"phone"`
	Created_at string `json:"created_at"`
	Deleted_at string `json:"deleted_at"`
	Is_admin   bool   `json:"is_admin"`
}

type SignupUser struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
