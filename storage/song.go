package storage

import (
	"MusicLibrary/models"
	"github.com/jmoiron/sqlx"
)

type SongStorage struct {
	db *sqlx.DB
}

func NewSongStorage(db *sqlx.DB) *SongStorage {
	return &SongStorage{
		db: db,
	}
}

func (s *SongStorage) GetSongs(filter, limit, offset string) ([]models.Song, error) {
	var songs []models.Song
	query := "SELECT * FROM songs WHERE group_name LIKE $1 LIMIT $2 OFFSET $3"
	err := s.db.Select(&songs, query, "%"+filter+"%", limit, offset)
	return songs, err
}

func (s *SongStorage) SaveSong(song *models.Song) error {
	query := "INSERT INTO songs (group_name, song_title, release_date, text, link) VALUES ($1, $2, $3, $4, $5)"
	_, err := s.db.Exec(query, song.GroupName, song.SongTitle, song.ReleaseDate, song.Text, song.Link)
	return err
}
