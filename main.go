package main

import (
	"log"
	"net/http"
	"strconv"

	"examples/htmx-go/components"
	"examples/htmx-go/models"

	"github.com/a-h/templ"
)

func main() {
	InitDatabase()
	defer DB.Close()

	mux := http.NewServeMux()
	mux.Handle("GET /", loggingMiddleware(http.HandlerFunc(indexHandler)))
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("POST /todos", loggingMiddleware(http.HandlerFunc(createToDoHandler)))
	mux.Handle("DELETE /todos/{id}", loggingMiddleware(http.HandlerFunc(deleteToDoHandler)))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	todos := ReadToDoList()
	render(w, r, components.Index(todos))
}

func createToDoHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	status := r.FormValue("status")
	id, _ := CreateToDo(title, status)
	render(w, r, components.Task(models.ToDo{Id: id, Title: title, Status: status}))
}

func deleteToDoHandler(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("id")
	id, _ := strconv.ParseInt(param, 10, 64)
	DeleteToDo(id)
}

func render(w http.ResponseWriter, r *http.Request, component templ.Component) error {
	return component.Render(r.Context(), w)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
