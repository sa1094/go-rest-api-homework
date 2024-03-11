package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sa1094/go-rest-api-homework/internal/controllers"
	"github.com/sa1094/go-rest-api-homework/internal/database"
	"github.com/sa1094/go-rest-api-homework/internal/models"
)

func initData() {
	var tasks = map[string]models.Task{
		"1": {
			ID:          "1",
			Description: "Сделать финальное задание темы REST API",
			Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
			Applications: []string{
				"VS Code",
				"Terminal",
				"git",
			},
		},
		"2": {
			ID:          "2",
			Description: "Протестировать финальное задание с помощью Postmen",
			Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
			Applications: []string{
				"VS Code",
				"Terminal",
				"git",
				"Postman",
			},
		},
	}

	db := database.GetInstance()

	for _, t := range tasks {
		db[t.ID] = t
	}

}

func main() {
	initData()
	r := chi.NewRouter()
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Get("/tasks", controllers.ListTasks)
	r.Post("/tasks", controllers.CreateTask)
	r.Get("/tasks/{id}", controllers.GetTaskByID)
	r.Delete("/tasks/{id}", controllers.DeleteTask)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
