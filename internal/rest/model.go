package rest

import (
	"News-portal/internal/newsportal"
	"net/url"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
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
		q.NewsId, q.CategoryId, q.TagId,
		q.PageSize, q.Page,
	)
}

func newErrorResponse(c echo.Context, statusCode int, errR error) error {
	return c.JSON(statusCode, errorResponse{errR.Error()})
}

func getParams(e url.Values) queryParams {
	newsID, _ := strconv.Atoi(e.Get("newsId"))
	categoryID, _ := strconv.Atoi(e.Get("categoryId"))
	tagID, _ := strconv.Atoi(e.Get("tagId"))
	pageSize, _ := strconv.Atoi(e.Get("pageSize"))
	page, _ := strconv.Atoi(e.Get("page"))

	//newsID, _ := strconv.Atoi(e.QueryParam("newsId"))
	//categoryID, _ := strconv.Atoi(e.QueryParam("categoryId"))
	//tagID, _ := strconv.Atoi(e.QueryParam("tagId"))
	//pageSize, _ := strconv.Atoi(e.QueryParam("pageSize"))
	//page, _ := strconv.Atoi(e.QueryParam("page"))

	return queryParams{
		NewsId:     newsID,
		CategoryId: categoryID,
		TagId:      tagID,
		PageSize:   pageSize,
		Page:       page,
	}
}
