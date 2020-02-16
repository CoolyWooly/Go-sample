package model

import (
	"github.com/jinzhu/gorm"
)

type ImageModel struct {
	gorm.Model
	IdExhibit uint   `json:"id_exhibit"`
	Url       string `json:"url"`
}
