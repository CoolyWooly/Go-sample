package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"sample_rest/model"
)

func GetAllTasks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var taskModels []model.TaskModel
	db.Find(&taskModels)
	respondJSON(w, http.StatusOK, taskModels)
}

func CreateTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	taskModel := model.TaskModel{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&taskModel); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&taskModel).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, taskModel)
}

func GetTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	taskModel := getTaskOr404(db, uint(idInt), w, r)
	if taskModel == nil {
		return
	}
	respondJSON(w, http.StatusOK, taskModel)
}

func UpdateTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	taskModel := getTaskOr404(db, uint(idInt), w, r)
	if taskModel == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&taskModel); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&taskModel).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, taskModel)
}

func DeleteTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	taskModel := getTaskOr404(db, uint(idInt), w, r)
	if taskModel == nil {
		return
	}
	if err := db.Delete(&taskModel).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisableTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	taskModel := getTaskOr404(db, uint(idInt), w, r)
	if taskModel == nil {
		return
	}
	taskModel.Disable()
	if err := db.Save(&taskModel).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, taskModel)
}

func EnableTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	idInt, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	taskModel := getTaskOr404(db, uint(idInt), w, r)
	if taskModel == nil {
		return
	}
	taskModel.Enable()
	if err := db.Save(&taskModel).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, taskModel)
}

// getTaskOr404 gets a task instance if exists, or respond the 404 error otherwise
func getTaskOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *model.TaskModel {
	taskModel := model.TaskModel{}
	if err := db.First(&taskModel, gorm.Model{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &taskModel
}
