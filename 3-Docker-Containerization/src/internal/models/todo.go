package models

import "github.com/google/uuid"

type Todo struct {
	Id          uuid.UUID
	Title       string
	Description string
	Priority    string
}

type TodoResponse struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"msg"`
	Description string    `json:"description,omitempty"`
	Priority    string    `json:"priority"`
}
