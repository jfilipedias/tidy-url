package web

import "github.com/labstack/echo/v4"

func home(c echo.Context) error {
	return c.String(200, "Hello World!")
}

func createURL(c echo.Context) error {
	return nil
}

func createURLPost(c echo.Context) error {
	return nil
}

func retrieveOriginalURL(c echo.Context) error {
	return nil
}
