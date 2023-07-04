package task

import (
	model "alg_app/internal/storage/model/task"
	"database/sql"
	"fmt"
	"log/slog"
	"strconv"
)

func AddTask(task *model.Task) error {
	const op = "storage.postgres.task.AddTask"

	db, err := sql.Open("postgres", "host=localhost port=5431 user=postgres password=Admin123 dbname=postgres sslmode=disable")
	if err != nil {
		slog.Error("Проблема с подключением ", slog.String("db", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}
	defer db.Close()

	var sumAmount float64
	err = db.QueryRow(`SELECT SUM(amount) FROM users WHERE login = $1`, task.Login).Scan(&sumAmount)
	if err != nil {
		slog.Error("Проблема с получением долга ", slog.String("db", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	if sumAmount+task.Amount < 1000 {
		if _, err = db.Exec(`INSERT INTO task(login, task, amount) values ($1,$2,$3)`, task.Login, task.Task, task.Amount); err != nil {
			slog.Error("Проблема с добавлением задачи ", slog.String("db", err.Error()))
			return fmt.Errorf("%s: %w", op, err)
		}

		if _, err := db.Exec(`UPDATE users SET amount = $1 WHERE login = $2`, sumAmount+task.Amount, task.Login); err != nil {
			slog.Error("Проблема с добавлением долга", slog.String("db", err.Error()))
			return fmt.Errorf("%s: %w", op, err)
		}
	} else {
		slog.Error("Сумма превышает лимит!")
		return fmt.Errorf("%s: %s", op, "сумма превышает лимит")
	}

	return nil
}

// CompleteTask - Функция вызывается при ответе на задачу и закрывает её устанавливая статус Resolved = true
func CompleteTask(taskId string, response string) error {
	const op = "storage.postgres.task.AddTask"
	id, _ := strconv.Atoi(taskId)

	db, err := sql.Open("postgres", "host=localhost port=5431 user=postgres password=Admin123 dbname=postgres sslmode=disable")
	if err != nil {
		slog.Error("Проблема с подключением ", slog.String("db", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}
	defer db.Close()

	_, err = db.Exec(`UPDATE task SET response = $1, resolved = true WHERE task_id = $2`, response, id)
	if err != nil {
		slog.Error("Проблема с отправкой ответа ", slog.String("db", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func Task(taskId string) (*model.Task, error) {
	const op = "storage.postgres.task.Task"
	id, _ := strconv.Atoi(taskId)

	db, err := sql.Open("postgres", "host=localhost port=5431 user=postgres password=Admin123 dbname=postgres sslmode=disable")
	if err != nil {
		slog.Error("Проблема с подключением ", slog.String("db", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer db.Close()

	var task model.Task
	err = db.QueryRow(`SELECT * FROM task WHERE task_id = $1`, id).Scan(
		&task.TaskId, &task.Login, &task.Task, &task.Amount, &task.Resolved, &task.Response)
	if err != nil {
		slog.Error("Проблема с получением задачи ", slog.String("db", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &task, nil
}
