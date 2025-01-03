package main

import (
	"github.com/cleanupDev/blog/handlers"
	"github.com/cleanupDev/blog/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	articleService := services.NewArticleService()
	homeHandler := handlers.NewHomeHandler(articleService)

	e.GET("/", homeHandler.HandleHome)

	e.Logger.Fatal(e.Start(":3000"))
}
