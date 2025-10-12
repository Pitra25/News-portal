package rest

import (
	"News-portal/internal/newsportal"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	queryParams struct {
		CategoryId int `form:"categoryId"`
		TagId      int `form:"tagId"`
		PageSize   int `form:"pageSize"`
		Page       int `form:"page"`
	}

	Tag struct {
		ID    int    `json:"tagID"`
		Title string `json:"title"`
	}

	Category struct {
		ID    int    `json:"categoryId"`
		Title string `json:"title"`
	}

	News struct {
		ID          int       `json:"newsID"`
		Title       string    `json:"title"`
		Content     *string   `json:"content"`
		Author      string    `json:"author"`
		Category    *Category `json:"category"`
		Tags        []Tag     `json:"tagIds"`
		PublishedAt time.Time `json:"publishedAt"`
	}

	NewsSummary struct {
		ID          int       `json:"newsSummaryID"`
		Title       string    `json:"title"`
		PublishedAt time.Time `json:"publishedAt"`
		Category    *Category `json:"category"`
		TagIds      []Tag     `json:"tagIds"`
	}

	errorResponse struct {
		Message string `json:"message"`
	}
)

func (q *queryParams) NewFilter() newsportal.Filters {
	return newsportal.NewFilters(
		q.CategoryId, q.TagId,
		q.PageSize, q.Page,
	)
}

func newErrorResponse(c echo.Context, statusCode int, errR error) error {
	return c.JSON(statusCode, errorResponse{errR.Error()})
}
