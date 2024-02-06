package controllers

import (
	"Appointy/Tasks/todo/db"
	"Appointy/Tasks/todo/structs"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetTasks(h *db.DBhandler, w http.ResponseWriter, r *http.Request) {
	var tasks []structs.Tasks
	h.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func PostTasks(h *db.DBhandler, w http.ResponseWriter, r *http.Request) {
	var temp structs.TempTasks
	json.NewDecoder(r.Body).Decode(&temp)
	var res = structs.Tasks{
		Task:   temp.Task,
		Time:   temp.Time,
		Status: temp.Status,
	}

	h.DB.Create(&res)
	json.NewEncoder(w).Encode(res)
}
func PutTasks(h *db.DBhandler, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/task/"))
	var temp structs.TempTasks
	json.NewDecoder(r.Body).Decode(&temp)
	h.DB.Save(&structs.Tasks{ID: id, Task: temp.Task, Time: temp.Time, Status: temp.Status})
	w.Write([]byte("updated"))
}

func DeleteTasks(h *db.DBhandler, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/task/"))
	h.DB.Delete(&structs.Tasks{}, id)
	w.Write([]byte("deleted"))
}

func PatchTasks(h *db.DBhandler, w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/task/"))
	var temp structs.TempTasks
	json.NewDecoder(r.Body).Decode(&temp)
	h.DB.Model(&structs.Tasks{}).Where("id", id).Update("status", temp.Status)
	w.Write([]byte("patched"))
}
