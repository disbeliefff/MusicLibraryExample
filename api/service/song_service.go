package service

import (
	"MusicLibrary/models"
	"MusicLibrary/storage"
)

type SongService struct {
	SongStorage    *storage.SongStorage
	ExternalAPIURL string
}

func NewSongService(storage *storage.SongStorage, apiURL string) *SongService {
	return &SongService{
		SongStorage:    storage,
		ExternalAPIURL: apiURL,
	}
}

func (s *SongService) GetSongs(filter, limit, offset string) ([]models.Song, error) {
	return s.SongStorage.GetSongs(filter, limit, offset)
}

func (s *SongService) GetSongText(id int, verseLimit string, verseOffset string) (string, error) {
	return s.SongStorage.GeSongText(id, verseLimit, verseOffset)
}

func (s *SongService) SaveSong(song *models.Song) error {
	return s.SongStorage.SaveSong(song)
}

func (s *SongService) UpdateSong(id int, song *models.Song) error {
	return s.SongStorage.UpdateSong(id, song)
}

func (s *SongService) DeleteSong(id int) error {
	return s.SongStorage.DeleteSong(id)
}
