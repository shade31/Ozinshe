package repository

type User struct {
	Id         int    `json:"id"`
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

// Создание юзера DB
// func InsertQuery(p *pgxpool.Pool, user SignupUser) error {
// 	_, err := p.Exec(context.Background(), "insert into users(email, password) values($1, $2) returning id", user.Email, user.Password)
// 	if err != nil {
// 		log.Fatal("Ошибка при записи данных в таблицу", err)
// 		return err
// 	}
// 	fmt.Println("Новая запись успешно добавлена в таблицу users")
// 	return nil
// }
