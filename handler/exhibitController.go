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
	var exhibitModels []model.ExhibitModel
	db.Find(&exhibitModels)
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

	if err := db.Save(&exhibitModel).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, exhibitModel)
}

func GetExhibit(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	taskModel := getExhibitOr404(db, uint(idInt), w, r)
	if taskModel == nil {
		return
	}
	respondJSON(w, http.StatusOK, taskModel)
}

func UpdateExhibit(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	exhibitModel := getExhibitOr404(db, uint(idInt), w, r)
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
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	exhibitModel := getExhibitOr404(db, uint(idInt), w, r)
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
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	exhibitModel := getExhibitOr404(db, uint(idInt), w, r)
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
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	exhibitModel := getExhibitOr404(db, uint(idInt), w, r)
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

func getExhibitOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *model.ExhibitModel {
	exhibitModel := model.ExhibitModel{}
	if err := db.First(&exhibitModel, gorm.Model{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	exhibitModel.Photos = getImagesById(db, id)
	return &exhibitModel
}

func getImagesById(db *gorm.DB, id uint) []string {
	images := []string{"", ""}
	var imageModels []model.ImageModel
	db.Find(&imageModels).Where("IdExhibit = ?", id)

	//imageModels.

	return images
}
