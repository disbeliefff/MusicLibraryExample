package models

import "time"

type Song struct {
	ID          int       `json:"id" db:"id"`
	GroupName   string    `db:"group_name"`
	SongTitle   string    `db:"song_title"`
	ReleaseDate string    `db:"release_date"`
	Text        string    `db:"text"`
	Link        string    `db:"link"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
