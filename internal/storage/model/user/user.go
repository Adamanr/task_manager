package user

type User struct {
	Login       string  `json:"login"`
	FIO         string  `json:"fio"`
	NumberGroup string  `json:"number_group"`
	Email       string  `json:"email"`
	Amount      float32 `json:"amount"`
}
