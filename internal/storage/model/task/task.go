package task

type Task struct {
	TaskId   int64   `json:"taskId"`
	Login    string  `json:"login"`
	Task     string  `json:"task"`
	Amount   float64 `json:"amount"`
	Resolved bool    `json:"resolved,omitempty"`
	Response string  `json:"response"`
}
