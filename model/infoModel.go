package model

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type InfoModel struct {
	Info string `json:"info"`
}
