package handler

import (
	"News-portal/internal/service"
	"log/slog"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	slog.Debug("handler initialization")
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
			news.GET("/", h.GetAllNews)
			news.GET("/short", h.GetAllShortNews)
			news.GET("/?categoryId=&tagId=&pageSize=&page=", h.GetAllNewsByQuery)
			news.GET("/:id", h.GetNewsById)

			count := news.Group("/count")
			{
				count.GET("/", h.GetNewsCount)
			}
		}

		tags := api.Group("/tags")
		{
			tags.GET("/", h.GetAllTags)
			tags.GET("/:id", h.GetTagById)
		}

		categories := api.Group("/categories")
		{
			categories.GET("/", h.GetAllCategories)
			categories.GET("/:id", h.GetCategoryById)
		}
	}

	return router
}
