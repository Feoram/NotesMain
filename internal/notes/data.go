package notes

import (
	"database/sql"
	"time"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) CreateNewNote(title, body string) error {
	_, err := r.db.Exec(`INSERT INTO note (title, body, created_at, user_id) VALUES ($1, $2, $3, $4)`, title, body, time.Now(), 1)
	if err != nil {
		return err
	}

	return nil
}
