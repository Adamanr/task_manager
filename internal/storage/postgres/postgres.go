package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
)

// Init - функция создаёт таблицы Users/Task при их отсутствии
func Init() error {
	const op = "storage.postgres.init"

	db, err := sql.Open("postgres", "host=localhost port=5431 user=postgres password=Admin123 dbname=postgres sslmode=disable")
	if err != nil {
		slog.Error("Проблема с подключением к дб", slog.String("db", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, err := db.Exec(`SELECT * FROM users LIMIT 1`); err != nil {
		_, err = db.Exec(`
			CREATE TABLE Users (
				Login VARCHAR(50) PRIMARY KEY  UNIQUE,
				FIO VARCHAR(100),
				Number_Group VARCHAR(10),
				Amount numeric DEFAULT 0.00,
				Email VARCHAR(50) UNIQUE
			);`)
		if err != nil {
			slog.Info("Проблема в создании таблицы users", slog.String("db", err.Error()))
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if _, err := db.Exec(`SELECT * FROM task LIMIT 1`); err != nil {
		_, err = db.Exec(`
			CREATE TABLE Task (
				Task_id SERIAL PRIMARY KEY UNIQUE,
				Login VARCHAR(50),
				Task VARCHAR(1500),
				Amount numeric DEFAULT 0.00,
				Resolved boolean default false,
				Response varchar(1500) default '',
				FOREIGN KEY (Login) REFERENCES "users" (Login)
			);`)
		if err != nil {
			slog.Info("Проблема в создании таблицы Task", slog.String("db", err.Error()))
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil
}
