package handlers

import (
	"net/http"
    "strconv"

	"examples/htmx-go/components"
	"examples/htmx-go/models"
)

type ToDoHandler struct {
    Todos models.ToDoModel 
}

func (h *ToDoHandler) Index(w http.ResponseWriter, r *http.Request) {
	todos := h.Todos.ReadToDoList()
	components.Render(w, r, components.Index(todos))
}

func (h *ToDoHandler) Create(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	status := r.FormValue("status")
	id, _ := h.Todos.Create(title, status)
    
    w.WriteHeader(http.StatusCreated)

	components.Render(w, r, components.Task(models.ToDo{Id: id, Title: title, Status: status}))
}

func (h *ToDoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("id")
	id, _ := strconv.ParseInt(param, 10, 64)
	h.Todos.Delete(id)
}
