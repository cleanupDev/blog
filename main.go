package main

import (
	"github.com/cleanupDev/blog/internal"
	"github.com/cleanupDev/blog/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var articles = []internal.Article{
	{ID: 1, Title: "Hello World", Summary: "My first post!", Content: "Full content..."},
	{ID: 2, Title: "Another Post", Summary: "Check it out!", Content: "More details..."},
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		component := pages.HomePage(articles)
		return component.Render(c.Request().Context(), c.Response().Writer)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
