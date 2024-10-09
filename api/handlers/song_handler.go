package handlers

import (
	"MusicLibrary/api/service"
	"MusicLibrary/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type SongHandler struct {
	SongService *service.SongService
}

func NewSongHandler(songService *service.SongService) *SongHandler {
	return &SongHandler{
		SongService: songService,
	}
}

func (h *SongHandler) GetSongs(c echo.Context) error {
	filter := c.QueryParam("filter")
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	songs, err := h.SongService.GetSongs(filter, limit, offset)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, songs)
}

func (h *SongHandler) GetSongText(c echo.Context) error {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid ID format"})
	}

	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	text, err := h.SongService.GetSongText(id, limit, offset)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, text)
}

func (h *SongHandler) SaveSong(c echo.Context) error {
	song := &models.Song{}
	err := c.Bind(song)
	if err != nil {
		return c.JSON(500, err)
	}
	err = h.SongService.SaveSong(song)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, song)
}

func (h *SongHandler) UpdateSong(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var song models.Song
	if err := c.Bind(&song); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.SongService.UpdateSong(id, &song); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, song)
}

func (h *SongHandler) DeleteSong(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(400, echo.Map{"error": "Invalid ID format"})
	}
	err = h.SongService.DeleteSong(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, nil)
}
