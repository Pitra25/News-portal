package rest

import (
	"News-portal/internal/newsportal"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *newsportal.Service
}

func New(service *newsportal.Service) *Handler {
	slog.Debug("rest initialization")
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		news := api.Group("/news")
		{
			news.GET("/", h.GetAllNews)           // +
			news.GET("/short", h.GetAllShortNews) // +
			news.GET("/:id", h.GetNewsById)       // +

			count := news.Group("/count")
			{
				count.GET("/", h.GetNewsCount) // +
			}
		}

		tags := api.Group("/tags")
		{
			tags.GET("/", h.GetAllTags)    // +
			tags.GET("/:id", h.GetTagById) // +
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", h.GetAllCategories)   // +
			categories.GET("/:id", h.GetCategoryById) // +
		}
	}

	return router
}
