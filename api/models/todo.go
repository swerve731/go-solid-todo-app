package models

import "github.com/google/uuid"

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	completed bool   `json:"completed"`
}

func CreateTodo(title string) Todo {
	u := uuid.New().String()

	t := Todo{
		ID:        u,
		Title:     title,
		completed: false,
	}

	return t
}

func (t *Todo) Complete() {
	t.completed = true
}
