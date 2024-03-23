package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// var Database *pgxpool.Pool

const (
	usersTable = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// func OpenDatabaseConnection() {
// 	var err error
// 	host := os.Getenv("POSTGRES_HOST")
// 	username := os.Getenv("POSTGRES_USER")
// 	password := os.Getenv("POSTGRES_PASSWORD")
// 	databaseName := os.Getenv("POSTGRES_DATABASE")
// 	port := os.Getenv("POSTGRES_PORT")

// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Douala", host, username, password, databaseName, port)

// 	// Create database connection
// 	pool, err := pgxpool.New(context.Background(), dsn)
// 	if err != nil {
// 		log.Fatal("Ошибка при получении соединения из пула БД")
// 	}

// 	Database = pool

// 	fmt.Println("Соединение к БД установлено!")
// }

// func CloseDatabaseConnection() {
// 	if Database != nil {
// 		Database.Close()
// 		fmt.Println("Соединение с БД закрыто.")
// 	}
// }
