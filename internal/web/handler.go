package web

import "github.com/labstack/echo/v4"

func home(c echo.Context) error {
	c.String(200, "Hello World!")
	return nil
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
