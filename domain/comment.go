package domain

import (
	"time"
)

type Comment struct {
	ID      string    `json:"_id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	MovieID string    `json:"movie_id"`
	Text    string    `json:"text"`
	Date    time.Time `json:"date"`
}
