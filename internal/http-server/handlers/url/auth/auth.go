package auth

import (
	user2 "alg_app/internal/storage/postgres/user"
	"github.com/gin-gonic/gin"
)

type User struct {
	Login       string
	FIO         string `json:"fio"`
	NumberGroup string `json:"number_group"`
	Email       string `json:"email"`
}

// CreateUser - Создаёт пользователя через поля FIO / NumberGroup / Email
func CreateUser(c *gin.Context) {
	var user *User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(500, gin.H{
			"message": "Data field not valid!",
			"error":   err.Error(),
		})
		return
	}

	err := user2.CreateUser(user.Email, user.FIO, user.NumberGroup)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Создание пользователя прошло безуспешно!",
			"error":   err.Error(),
		})
		return
	}

	c.Status(200)
	return
}

// GetUser - Получает пользователя через параметр логина /user/nikolay
func GetUser(c *gin.Context) {
	login := c.Param("login")

	users, err := user2.GetUser(login)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Пользователя не получилось достатб..",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"user": users})
}
