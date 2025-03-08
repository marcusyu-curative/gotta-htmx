package main

import (
	"fmt"
	"net/http"

	"examples/htmx-go/db"
	"examples/htmx-go/handlers"
	"examples/htmx-go/middleware"
	"examples/htmx-go/models"
)

func main() {
    database, _ := db.New()
	defer database.Close()
    
	router := http.NewServeMux()
	router.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
    todoHandler := &handlers.ToDoHandler{
        Todos: models.ToDoModel{DB: database.Conn},
    }

    router.HandleFunc("GET /", todoHandler.Index)
    router.HandleFunc("POST /todos", todoHandler.Create) 
    router.HandleFunc("DELETE /todos/{id}", todoHandler.Delete)

    invoiceHandler := &handlers.InvoiceHandler{}

    adminRouter := http.NewServeMux()
    adminRouter.HandleFunc("POST /invoice", invoiceHandler.Create)

    // Middleware stacking
    stack := middleware.CreateStack(
        middleware.Logging,
    )

    // Subrouting
    v1 := http.NewServeMux()
    v1.Handle("/v1/", http.StripPrefix("/v1", router))

    adminStack := middleware.CreateStack(
        middleware.EnsureAdmin,
        middleware.IsAuthenticated,
    )
    router.Handle("/", adminStack(adminRouter))
    
    server := http.Server{
        Addr:    ":8080",
        Handler: stack(router),
    }

    fmt.Println("Server listening port 8080")
    server.ListenAndServe()
}
