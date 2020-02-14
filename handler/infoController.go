package handler

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"sample_rest/model"
)

func GetInfo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	infoModel := model.InfoModel{Info: "Приложение Digital Art открывает для вас все мировые экспонаты и выставки"}
	respondJSON(w, http.StatusOK, infoModel)
}
