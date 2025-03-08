package middleware

import (
    "log"
    "net/http"
    "strings"
)

func EnsureAdmin(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Checking if user is an admin")
        if !strings.Contains(r.Header.Get("UserType"), "Admin") {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}
            
