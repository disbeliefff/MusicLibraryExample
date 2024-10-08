package routers

import "github.com/labstack/echo/v4"

func SetupRouter(e *echo.Echo) {
	e.GET("/songs", func(c echo.Context) error {
		return c.String(200, "List of songs")
	})
}
