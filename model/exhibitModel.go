package model

import (
	"github.com/jinzhu/gorm"
)

type ExhibitModel struct {
	gorm.Model
	MuseumId    string   `json:"museum_id"`
	Rating      float32  `json:"rating"`
	Name        int      `json:"name"`
	Description string   `json:"description"`
	Year        string   `json:"year"`
	Author      string   `json:"author"`
	Audio       string   `json:"audio"`
	Photos      []string `json:"photos"`
	IsPopular   bool     `json:"is_popular"`
}

func (e *ExhibitModel) RemoveFromPopular() {
	e.IsPopular = false
}

func (e *ExhibitModel) AddToPopular() {
	e.IsPopular = true
}
