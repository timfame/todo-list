package http

import (
	"github.com/timfame/todo-list/internal/models"
	"log"
	"net/http"
)

func (s *server) index(w http.ResponseWriter, r *http.Request) {
	lists, err := s.service.GetLists(r.Context())
	if err != nil {
		log.Println("Get lists error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := s.tmpl.Execute(w, lists); err != nil {
		log.Println("Execute error:", err)
	}
}

func (s *server) addList(w http.ResponseWriter, r *http.Request) {
	list := models.List{
		Title: r.PostFormValue("title"),
	}
	if err := s.service.AddList(r.Context(), list); err != nil {
		log.Println("Add list error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) removeList(w http.ResponseWriter, r *http.Request) {
	listIndex, err := getPostFormInt(w, r, "listIndex")
	if err != nil {
		return
	}
	if err := s.service.RemoveList(r.Context(), listIndex); err != nil {
		log.Println("Remove list error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) addTask(w http.ResponseWriter, r *http.Request) {
	listIndex, err := getPostFormInt(w, r, "listIndex")
	if err != nil {
		return
	}
	task := models.Task{
		Content: r.PostFormValue("content"),
	}
	if err := s.service.AddTask(r.Context(), listIndex, task); err != nil {
		log.Println("Add task error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (s *server) markTaskDone(w http.ResponseWriter, r *http.Request) {
	listIndex, err := getPostFormInt(w, r, "listIndex")
	if err != nil {
		return
	}
	taskIndex, err := getPostFormInt(w, r, "taskIndex")
	if err != nil {
		return
	}
	if err := s.service.MarkTaskDone(r.Context(), listIndex, taskIndex); err != nil {
		log.Println("Mark task done error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
