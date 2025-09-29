package models

import (
	"time"
)

const (
	NewsTable       = "news"
	StatusesTable   = "statuses"
	TagsTable       = "tags"
	CategoriesTable = "categories"

	NewsStatus = 1
)

// Table DB
type (
	News struct {
		NewsID      int       `db:"newsId" json:"newsId"`
		Title       string    `db:"title" json:"title"`
		Content     string    `db:"content" json:"content"`
		Author      string    `db:"author" json:"author"`
		CategoryID  int       `db:"categoryId" json:"categoryId"`
		TagIds      []int     `db:"tagIds" json:"tagIds"`
		CreatedAt   time.Time `db:"createdAt,readonly" json:"createdAt"`
		PublishedAt time.Time `db:"publishedAt" json:"publishedAt"`
		StatusID    int       `db:"statusId" json:"statusId"`
	}

	ShortNews struct {
		NewsID      int       `db:"newsId" json:"newsId"`
		Title       string    `db:"title" json:"title"`
		PublishedAt time.Time `db:"publishedAt" json:"publishedAt"`
		CategoryID  int       `db:"categoryId" json:"categoryId"`
		TagIds      []int     `db:"tagIds" json:"tagIds"`
	}

	Categories struct {
		CategoryID  int    `db:"categoryId" json:"categoryId"`
		Title       string `db:"title" json:"title"`
		OrderNumber int    `db:"orderNumber" json:"orderNumber"`
		StatusID    int    `db:"statusId" json:"statusId"`
	}

	Tags struct {
		TagID    int    `db:"tagId" json:"tagId"`
		Title    string `db:"title" json:"title"`
		StatusID int    `db:"statusId" json:"statusId"`
	}

	Statuses struct {
		StatusID int    `db:"statusId" json:"statusId"`
		Title    string `db:"title" json:"title"`
	}
)

// Response
type (
	ResponseStatuses struct {
		StatusID int    `json:"statusId"`
		Title    string `json:"title"`
	}

	ResponseTag struct {
		TagID int    `json:"tagId"`
		Title string `json:"title"`
	}

	ResponseNews struct {
		NewsID      int              `json:"newsId"`
		Title       string           `json:"title"`
		Content     string           `json:"content"`
		Author      string           `json:"author"`
		CategoryID  string           `json:"categoryId"`
		TagIds      []ResponseTag    `json:"tagIds"`
		CreatedAt   time.Time        `json:"createdAt"`
		PublishedAt time.Time        `json:"publishedAt"`
		StatusID    ResponseStatuses `json:"statusId"`
	}

	ResponseShortNews struct {
		NewsID      int           `json:"newsId"`
		Title       string        `json:"title"`
		Content     string        `json:"content"`
		PublishedAt time.Time     `json:"publishedAt"`
		CategoryID  string        `json:"categoryId"`
		TagIds      []ResponseTag `json:"tagIds"`
	}

	ResponseCategories struct {
		CategoryID  int              `json:"categoryId"`
		Title       string           `json:"title"`
		OrderNumber int              `json:"orderNumber"`
		StatusID    ResponseStatuses `json:"statusId"`
	}

	ResponseTags struct {
		TagID    int              `json:"tagId"`
		Title    string           `json:"title"`
		StatusID ResponseStatuses `json:"statusId"`
	}
)
