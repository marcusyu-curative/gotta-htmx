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

	// e := gin.Default()
	// e.LoadHTMLGlob("templates/*")
	mux := http.NewServeMux()

	// e.GET("/", func(c *gin.Context) {
	// 	todos := ReadToDoList()
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"todos": todos,
	// 	})
	// })
	mux.Handle("/", loggingMiddleware(http.HandlerFunc(indexHandler)))

	// e.POST("/todos", func(c *gin.Context) {
	// 	title := c.PostForm("title")
	// 	status := c.PostForm("status")
	// 	id, _ := CreateToDo(title, status)

	/* c.HTML leverages Go native templating */
	// c.HTML(http.StatusOK, "task", gin.H{
	// 	"Title":  title,
	// 	"Status": status,
	// 	"Id":     id,
	// })

	// 	c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	// 	c.Status(http.StatusOK)
	// 	component := components.Task(models.ToDo{Id: id, Title: title, Status: status})
	// 	if err := component.Render(c.Request.Context(), c.Writer); err != nil {
	// 		c.String(http.StatusInternalServerError, err.Error())
	// 	}
	// })
	mux.Handle("POST /todos", loggingMiddleware(http.HandlerFunc(createToDoHandler)))

	// e.DELETE("/todos/:id", func(c *gin.Context) {
	// 	param := c.Param("id")
	// 	id, _ := strconv.ParseInt(param, 10, 64)
	// 	DeleteToDo(id)
	// })
	mux.Handle("DELETE /todos/{id}", loggingMiddleware(http.HandlerFunc(deleteToDoHandler)))

	//e.Run(":8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	todos := ReadToDoList()
	// if err := comp.Render(r.Context(), w); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
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
