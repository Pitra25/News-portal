package rest

import (
	"News-portal/internal/newsportal"

	"github.com/labstack/echo/v4"
)

type Router struct {
	manager *newsportal.Manager
}

func NewRouter(manager *newsportal.Manager) *Router {
	return &Router{
		manager: manager,
	}
}

func (h *Router) AddRouter(e *echo.Echo) {
	api := e.Group("/api")

	news := api.Group("/news")
	news.GET("/", h.GetAllNews)
	news.GET("/short", h.GetNewsSummaries)
	news.GET("/:id", h.GetNewsById)

	count := news.Group("/count")
	count.GET("/", h.GetNewsCount)

	tags := api.Group("/tags")
	tags.GET("/", h.GetAllTags)

	categories := api.Group("/categories")
	categories.GET("/", h.GetAllCategories)

	return
}
