package main

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.GET("/", homeHandler)

	e.GET("/test-htmx", func(c echo.Context) error {
		return c.String(http.StatusOK, "<p>Dynamic content loaded via HTMX!</p>")
	})

	e.Start(":8042")
}

func homeHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := map[string]interface{}{
		"Title": "simple blog",
	}
	return tmpl.Execute(c.Response().Writer, data)
}
