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
	news.POST("/", h.AddNews)
	news.PUT("/", h.UpdateNews)
	news.DELETE("/:id", h.DeleteNews)

	count := news.Group("/count")
	count.GET("/", h.GetNewsCount)

	tags := api.Group("/tags")
	tags.GET("/", h.GetAllTags)
	tags.POST("/", h.AddTag)
	tags.PUT("/", h.UpdateTag)
	tags.DELETE("/:id", h.DeleteTag)

	categories := api.Group("/categories")
	categories.GET("/", h.GetAllCategories)
	categories.POST("/", h.AddCategory)
	categories.PUT("/", h.UpdateCategory)
	categories.DELETE("/:id", h.DeleteCategory)

	return
}
