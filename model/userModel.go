package model

import (
	"github.com/jinzhu/gorm"
)

type UserModel struct {
	gorm.Model
	Name   string `gorm:"unique" json:"name"`
	City   string `json:"city"`
	Age    int    `json:"age"`
	Status bool   `json:"status"`
}
