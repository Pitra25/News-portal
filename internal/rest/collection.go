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

//go:generate colgen -imports=News-portal/internal/newsportal
//colgen:News,Category,Tag
//colgen:News:MapP(newsportal)
//colgen:Tag:MapP(newsportal)
//colgen:Category:MapP(newsportal)

// MapP converts slice of type T to slice of type M with given converter with pointers.
func MapP[T, M any](a []T, f func(*T) *M) []M {
	n := make([]M, len(a))
	for i := range a {
		n[i] = *f(&a[i])
	}
	return n
}

func NewCategory(in *newsportal.Category) *Category {
	if in == nil {
		return nil
	}

	return &Category{
		ID:    in.ID,
		Title: in.Title,
	}
}

func NewNews(in *newsportal.News) *News {
	if in == nil {
		return nil
	}

	return &News{
		ID:          in.ID,
		Title:       in.Title,
		Author:      in.Author,
		Content:     in.Content,
		PublishedAt: in.PublishedAt,
		Category:    NewCategory(in.Category),
	}
}

func NewTag(in *newsportal.Tag) *Tag {
	if in == nil {
		return nil
	}

	return &Tag{
		ID:    in.ID,
		Title: in.Title,
	}
}

func NewNewsSummaries(in []newsportal.News) []NewsSummary {
	if in == nil {
		return nil
	}
	result := make([]NewsSummary, len(in))
	for _, i := range in {
		result = append(result, *NewNewsSummary(&i))
	}
	return result
}

func NewNewsSummary(in *newsportal.News) *NewsSummary {
	if in == nil {
		return nil
	}
	return &NewsSummary{
		ID:          in.ID,
		Title:       in.Title,
		PublishedAt: in.PublishedAt,
		Category:    NewCategory(in.Category),
		TagIds:      NewTags(in.Tags),
	}
}
