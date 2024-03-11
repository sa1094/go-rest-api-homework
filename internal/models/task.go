package models

import (
	"fmt"

	"github.com/sa1094/go-rest-api-homework/internal/database"
)

type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

func (t *Task) Create() error {
	db := database.GetInstance()
	if _, ok := db[t.ID]; ok {
		return fmt.Errorf("item already exist")
	}
	db[t.ID] = *t
	return nil
}

func Delete(id string) error {
	db := database.GetInstance()
	if _, ok := db[id]; !ok {
		return fmt.Errorf("item doesn't exist")
	}
	delete(db, id)
	return nil
}

func List() []Task {
	result := []Task{}
	db := database.GetInstance()
	for _, v := range db {
		if t, ok := v.(Task); ok {
			result = append(result, t)
		}
	}
	return result
}

func ByID(id string) (Task, error) {
	db := database.GetInstance()
	if _, ok := db[id]; !ok {
		return Task{}, fmt.Errorf("item doesn't exist")
	}
	t, ok := db[id].(Task)
	if !ok {
		return Task{}, fmt.Errorf("item doesn't exist")
	}
	return t, nil

}
