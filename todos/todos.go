package todos

import "time"

type Todo struct {
	Id        uint      `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	Added     time.Time `json:"created_at"`
}

func NewTodo(id uint, text string) *Todo {
	return &Todo{
		Id:        id,
		Text:      text,
		Completed: false,
		Added:     time.Now(),
	}
}
