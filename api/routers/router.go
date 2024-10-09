package routers

import (
	"MusicLibrary/api/handlers"
	"MusicLibrary/api/service"
	"MusicLibrary/storage"
	"github.com/labstack/echo/v4"
)

func SetupRouter(e *echo.Echo, storage *storage.SongStorage) {
	songService := service.NewSongService(storage, "https://www.youtube.com/")
	h := handlers.NewSongHandler(songService)

	e.GET("/songs", h.GetSongs)
	e.GET("/songs/:id", h.GetSongText)
	e.POST("/songs", h.SaveSong)
	e.PUT("/songs/:id", h.UpdateSong)
	e.DELETE("/songs/:id", h.DeleteSong)
}
