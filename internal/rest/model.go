package rest

import (
	"News-portal/internal/newsportal"
	"time"
)

type (
	queryParams struct {
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
)

func (q *queryParams) NewFilter() newsportal.Filters {
	return newsportal.Filters{
		News: newsportal.NewsFilters{
			NewsId:     0,
			TagId:      q.TagId,
			CategoryId: q.CategoryId,
		},
		Page: newsportal.PageFilters{
			PageSize: q.PageSize,
			Page:     q.Page,
		},
	}
}
