package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func routes() http.Handler {
	e := echo.New()
	e.GET("/", home)
	e.GET("/url", createURL)
	e.POST("/url", createURLPost)
	e.GET("/:hash", retrieveOriginalURL)

	return e
}
