package storage

import (
	"MusicLibrary/models"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type SongStorage struct {
	db *sqlx.DB
}

func (s *SongStorage) GetSongs(filter, limit, offset string) ([]models.Song, error) {
	var songs []models.Song
	query := "SELECT * FROM songs WHERE group_name LIKE $1 LIMIT $2 OFFSET $3"
	err := s.db.Select(&songs, query, "%"+filter+"%", limit, offset)
	return songs, err
}

func (s *SongStorage) GeSongText(id int, verseLimit, verseOffset string) (string, error) {
	var text string
	query := `
		SELECT text 
		FROM songs 
		WHERE id = $1 
		OFFSET $2 LIMIT $3`
	err := s.db.Get(&text, query, id, verseOffset, verseLimit)
	return text, err
}

func (s *SongStorage) SaveSong(song *models.Song) error {
	query := "INSERT INTO songs (group_name, song_title, release_date, text, link) VALUES ($1, $2, $3, $4, $5)"
	_, err := s.db.Exec(query, song.GroupName, song.SongTitle, song.ReleaseDate, song.Text, song.Link)
	return err
}

func (s *SongStorage) DeleteSong(id int) error {
	query := "DELETE FROM songs WHERE id = $1"
	result, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return echo.NewHTTPError(404, "Song not found")
	}
	return nil
}

func (s *SongStorage) UpdateSong(id int, song *models.Song) error {
	query := `
		UPDATE songs 
		SET group_name = $1, song_title = $2, release_date = $3, text = $4, link = $5 
		WHERE id = $6`
	_, err := s.db.Exec(query, song.GroupName, song.SongTitle, song.ReleaseDate, song.Text, song.Link, id)
	return err
}
