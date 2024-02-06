package handler

import (
	"Appointy/Tasks/todo/controllers"
	"Appointy/Tasks/todo/db"
	"net/http"
)

func TasksHandler(h *db.DBhandler, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetTasks(h, w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func AddHandler(h *db.DBhandler, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controllers.PostTasks(h, w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ChangeHandler(h *db.DBhandler, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		controllers.PutTasks(h, w, r)
	case http.MethodDelete:
		controllers.DeleteTasks(h, w, r)
	case http.MethodPatch:
		controllers.PatchTasks(h, w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
