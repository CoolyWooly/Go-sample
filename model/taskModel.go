package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
