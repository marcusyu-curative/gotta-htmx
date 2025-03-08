package models

import "time"

type Invoice struct {
    Id uint `json:"id"`
    Name string `json:"name"`
    Amount uint `json:"amount"`
    Timestamp time.Time `json:"timestamp"`
}
