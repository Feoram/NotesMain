package notes

import "time"

type Note struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
