package rest

import (
	"News-portal/internal/newsportal"
	"net/http"
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

	updateStatus struct {
		NewsId     *int `json:"newsId"`
		CategoryId *int `json:"categoryId"`
		TagId      *int `json:"tagId"`
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

	NewsInput struct {
		Id          *int    `json:"newsID"`
		Title       string  `json:"title"`
		Content     *string `json:"content"`
		Author      string  `json:"author"`
		CategoryID  int     `json:"categoryID"`
		Tags        []int   `json:"tagIds"`
		PublishedAt string  `json:"publishedAt"`
	}

	TagInput struct {
		ID    *int   `json:"tagID"`
		Title string `json:"title"`
	}

	CategoryInput struct {
		ID          *int   `json:"categoryID"`
		Title       string `json:"title"`
		OrderNumber *int   `json:"orderNumber"`
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

func newNoContentResponse(c echo.Context, status int, res bool) error {
	if res {
		return c.NoContent(status)
	} else {
		return c.NoContent(http.StatusInternalServerError)
	}
}

func newsToManager(in *NewsInput) *newsportal.NewsInput {
	layout := "2006-01-02 15:04:05.000 -0700"
	timeP, _ := time.Parse(layout, in.PublishedAt)

	return &newsportal.NewsInput{
		Id:          in.Id,
		Title:       in.Title,
		Content:     in.Content,
		Author:      in.Author,
		CategoryID:  in.CategoryID,
		TagIDs:      in.Tags,
		PublishedAt: &timeP,
	}
}

func categoryToManager(in *CategoryInput) *newsportal.CategoryInput {
	return &newsportal.CategoryInput{
		ID:          in.ID,
		Title:       in.Title,
		OrderNumber: in.OrderNumber,
	}
}

func tagToManager(in *TagInput) *newsportal.TagInput {
	return &newsportal.TagInput{
		ID:    in.ID,
		Title: in.Title,
	}
}

func updateStatusToManager(in *updateStatus) newsportal.UpdateStatus {
	return newsportal.UpdateStatus{
		NewsId:     in.NewsId,
		CategoryId: in.CategoryId,
		TagId:      in.TagId,
	}
}
