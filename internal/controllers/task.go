package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sa1094/go-rest-api-homework/internal/models"
)

func CreateTask(res http.ResponseWriter, req *http.Request) {
	task := &models.Task{}
	body, err := io.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, task)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	err = task.Create()
	if err != nil {
		http.Error(res, err.Error(), http.StatusConflict)
		return
	}

}

func DeleteTask(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	err := models.Delete(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}

func ListTasks(res http.ResponseWriter, _ *http.Request) {
	r := models.List()
	s, err := json.Marshal(r)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	_, err = res.Write(s)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

}

func GetTaskByID(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	task, err := models.ByID(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	s, err := json.Marshal(task)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = res.Write(s)

	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
