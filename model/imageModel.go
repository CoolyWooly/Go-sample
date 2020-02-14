package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ImageModel struct {
	gorm.Model
	IdExhibit uint   `json:"id_exhibit"`
	Url       string `json:"url"`
}
