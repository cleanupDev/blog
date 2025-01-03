package handlers

import (
	"github.com/cleanupDev/blog/components"
	"github.com/cleanupDev/blog/services"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	articleService *services.ArticleService
}

func NewHomeHandler(articleService *services.ArticleService) *HomeHandler {
	return &HomeHandler{articleService: articleService}
}

func (h *HomeHandler) HandleHome(c echo.Context) error {
	articles := h.articleService.GetAllArticles()
	component := components.HomePage(articles)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
