package handlers

import (
	"examples/htmx-go/middleware"
	"log"
	"net/http"
)

type InvoiceHandler struct {
}

func (h *InvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
    userID, ok := r.Context().Value(middleware.AuthUserID).(string)
    if !ok {
        log.Println("Invalid user ID")
        w.WriteHeader(http.StatusBadRequest)
    }

    log.Println("Creating invoice for user:", userID)

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Invoice created for " + userID))
}

