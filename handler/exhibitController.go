package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"sample_rest/model"
)

func GetAllExhibits(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	page := r.FormValue("page")
	limit := r.FormValue("limit")
	search := r.FormValue("search")

	pageInt, _ := strconv.ParseInt(page, 10, 32)
	limitInt, _ := strconv.ParseInt(limit, 10, 32)
	var exhibitModels []model.ExhibitModel
	db.Offset((pageInt-1)*limitInt).Limit(limitInt).Preload("Images").Where("Name LIKE ?", search+"%").Find(&exhibitModels)

	respondJSON(w, http.StatusOK, exhibitModels)
}

func CreateExhibit(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	exhibitModel := model.ExhibitModel{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&exhibitModel); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Create(&exhibitModel).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, exhibitModel)
}

func GetExhibit(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	taskModel := getExhibitOr404(db, int(idInt), w, r)
	if taskModel == nil {
		return
	}
	respondJSON(w, http.StatusOK, taskModel)
}

func UpdateExhibit(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	exhibitModel := getExhibitOr404(db, int(idInt), w, r)
	if exhibitModel == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&exhibitModel); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&exhibitModel).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, exhibitModel)
}

func DeleteExhibit(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	exhibitModel := getExhibitOr404(db, int(idInt), w, r)
	if exhibitModel == nil {
		return
	}
	if err := db.Delete(&exhibitModel).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisableExhibit(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	exhibitModel := getExhibitOr404(db, int(idInt), w, r)
	if exhibitModel == nil {
		return
	}
	exhibitModel.RemoveFromPopular()
	if err := db.Save(&exhibitModel).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, exhibitModel)
}

func EnableExhibit(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	exhibitModel := getExhibitOr404(db, int(idInt), w, r)
	if exhibitModel == nil {
		return
	}
	exhibitModel.AddToPopular()
	if err := db.Save(&exhibitModel).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, exhibitModel)
}

func getExhibitOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.ExhibitModel {
	exhibitModel := model.ExhibitModel{}
	if err := db.Preload("Images").First(&exhibitModel, model.ExhibitModel{IdExhibit: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &exhibitModel
}
