package rest

import (
	"News-portal/internal/newsportal"

	"github.com/gin-gonic/gin"
)

type Router struct {
	manager *newsportal.Manager
}

func NewRouter(manager *newsportal.Manager) *Router {
	return &Router{
		manager: manager,
	}
}

func (h *Router) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		news := api.Group("/news")
		{
			news.GET("/", h.GetAllNews)
			news.GET("/short", h.GetAllShortNews)
			news.GET("/:id", h.GetNewsById)

			count := news.Group("/count")
			{
				count.GET("/", h.GetNewsCount)
			}
		}

		tags := api.Group("/tags")
		{
			tags.GET("/", h.GetAllTags)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", h.GetAllCategories)
		}
	}

	return router
}
