package user

import (
	model "alg_app/internal/storage/model/user"
	"database/sql"
	"fmt"
	"log/slog"
	"strings"
)

// CreateUser - создаёт пользователя в БД через поля Email / FIO / NumberGroup
func CreateUser(email, fio, numberGroup string) error {
	const op = "storage.postgres.user.CreateUser"
	login := strings.Split(email, "@")[0]

	db, err := sql.Open("postgres", "host=localhost port=5431 user=postgres password=Admin123 dbname=postgres sslmode=disable")
	if err != nil {
		slog.Error("Проблема с подключением ", slog.String("db", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}
	defer db.Close()

	var count int

	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE login = $1 LIMIT 1", login).Scan(&count)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if count > 0 {
		return fmt.Errorf("пользователь уже существует в базе данных")
	}

	if _, err = db.Exec(`INSERT INTO users(login, fio, number_group, email) VALUES ($1, $2, $3,$4)`, login, fio, numberGroup, email); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

// GetUser - Получает по логину из БД информацию о пользователе
func GetUser(login string) (*model.User, error) {
	const op = "storage.postgres.user.GetUser"

	db, err := sql.Open("postgres", "host=localhost port=5431 user=postgres password=Admin123 dbname=postgres sslmode=disable")
	if err != nil {
		slog.Error("Проблема с подключением ", slog.String("db", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer db.Close()

	var user model.User
	err = db.QueryRow("SELECT * FROM users WHERE Login = $1", login).Scan(&user.Login, &user.FIO, &user.NumberGroup, &user.Amount, &user.Email)
	if err != nil {
		slog.Error("Проблема c получением пользователя ", slog.String("db", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}
