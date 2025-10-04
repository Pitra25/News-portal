package rest

import (
	"News-portal/internal/newsportal"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	queryParams struct {
		NewsId     int `form:"newsId"`
		CategoryId int `form:"categoryId"`
		TagId      int `form:"tagId"`
		PageSize   int `form:"pageSize"`
		Page       int `form:"page"`
	}

	Tag struct {
		TagID int    `json:"tagId"`
		Title string `json:"title"`
	}

	Category struct {
		CategoryID int    `json:"categoryId"`
		Title      string `json:"title"`
	}
	News struct {
		NewsID      int       `json:"newsId"`
		Title       string    `json:"title"`
		Content     string    `json:"content"`
		Author      string    `json:"author"`
		Category    Category  `json:"category"`
		Tags        []Tag     `json:"tagIds"`
		PublishedAt time.Time `json:"publishedAt"`
	}

	ShortNews struct {
		NewsID      int       `json:"newsId"`
		Title       string    `json:"title"`
		PublishedAt time.Time `json:"publishedAt"`
		Category    Category  `json:"category"`
		TagIds      []Tag     `json:"tagIds"`
	}

	errorResponse struct {
		Message string `json:"message"`
	}
)

func (q *queryParams) NewFilter() newsportal.Filters {
	filter := newsportal.NewFilters(
		q.NewsId, q.CategoryId, q.TagId,
		q.PageSize, q.Page,
	)
	return filter
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
