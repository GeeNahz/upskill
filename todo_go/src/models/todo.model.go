package models

import (
	"gorm.io/gorm"
)

type TodoStatus string

const (
	STATUS_COMPLETED  TodoStatus = "COMPLETED"
	STATUS_INCOMPLETE TodoStatus = "INCOMPLETED"
)

type Todo struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
}
