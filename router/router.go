package router

import (
	// Import the missing package
	"Appointy/Tasks/todo/db"
	"Appointy/Tasks/todo/handler" // Import the missing package
	"net/http"
)

func CreateRouter(h db.DBhandler) *http.ServeMux {

	r := http.NewServeMux()
	r.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		handler.TasksHandler(&h, w, r)
	})
	r.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		handler.AddHandler(&h, w, r)
	})
	r.HandleFunc("/task/", func(w http.ResponseWriter, r *http.Request) {
		handler.ChangeHandler(&h, w, r)
	})
	return r
}
