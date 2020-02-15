package model

import (
	"github.com/jinzhu/gorm"
)

type TaskModel struct {
	gorm.Model
	Title   string `json:"title"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func (e *TaskModel) Disable() {
	e.Status = false
}

func (e *TaskModel) Enable() {
	e.Status = true
}
