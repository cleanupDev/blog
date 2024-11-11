package main

import (
	"html/template"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Post struct {
	ID    int
	Title string
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.GET("/", homeHandler)

	e.GET("/test-htmx", func(c echo.Context) error {
		return c.String(http.StatusOK, "<p>Dynamic content loaded via HTMX!</p>")
	})

	e.GET("/posts/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, "Post ID: "+id)
	})
	e.GET("/posts", postsHandler)

	e.Start(":3000")
}

func homeHandler(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := map[string]interface{}{
		"Title": "simple blog",
	}
	return tmpl.Execute(c.Response().Writer, data)
}

func postsHandler(c echo.Context) error {
	posts := []Post{
		{ID: 3, Title: "Post 3"},
		{ID: 1, Title: "Post 1"},
		{ID: 2, Title: "Post 2"},
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].ID < posts[j].ID
	})

	tmpl := template.Must(template.ParseFiles("templates/posts.html"))

	data := map[string]interface{}{
		"Posts": posts,
	}
	return tmpl.Execute(c.Response().Writer, data)
}
