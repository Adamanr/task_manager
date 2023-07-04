package main

import (
	"alg_app/internal/http-server/handlers/url/auth"
	"alg_app/internal/http-server/handlers/url/task"
	"alg_app/internal/storage/postgres"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
)

func main() {
	router := gin.New()

	if err := postgres.Init(); err != nil {
		log.Fatalf("Error init db: %v", err)
	}

	router.GET("/student/new", auth.CreateUser)
	router.GET("/student/:login", auth.GetUser)
	router.POST("/student/task/new", task.NewTask)
	router.POST("/student/task/:id", task.CompleteTask)
	router.GET("/student/task/:id", task.Task)

	if err := router.Run("localhost:8080"); err != nil {
		slog.Error("Error run server", slog.String("err", err.Error()))
		return
	}
}
