package task

import (
	model "alg_app/internal/storage/model/task"
	controller "alg_app/internal/storage/postgres/task"
	"github.com/gin-gonic/gin"
)

// NewTask - Создаёт задачу по полям Login / Task / Amount
func NewTask(c *gin.Context) {
	var task *model.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(500, gin.H{
			"message": "Поля не соответствуют",
			"error":   err.Error(),
		})
		return
	}

	err := controller.AddTask(task)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Ошибка добавления долга",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"Task": task})
}

// CompleteTask - Выполняет задачу получая id через параметр и поле response - которое является ответом на задачу
func CompleteTask(c *gin.Context) {
	req := make(map[string]string, 0)
	id := c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		c.JSON(500, gin.H{
			"message": "Поля не соответствуют",
			"error":   err.Error(),
		})
		return
	}

	if err := controller.CompleteTask(id, req["response"]); err != nil {
		c.JSON(500, gin.H{
			"message": "Проблема с завершением задачи",
			"error":   err.Error(),
		})
		return
	}

	c.Status(200)
}

// Task - получает таск по его id
func Task(c *gin.Context) {
	id := c.Param("id")

	task, err := controller.Task(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Проблема с завершением задачи",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"task": task})
}
