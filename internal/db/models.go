package db

import (
	"time"

	"github.com/lib/pq"
)

const (
	newsTable       = "newsportal.news"
	statusesTable   = "newsportal.statuses"
	tagsTable       = "newsportal.tags"
	categoriesTable = "newsportal.categories"

	newsStatus = 1
)

// Table DB
type (
	News struct {
		NewsID      int           `db:"newsId"`
		Title       string        `db:"title"`
		Content     string        `db:"content"`
		Author      string        `db:"author"`
		CategoryID  int           `db:"categoryId"`
		TagIds      pq.Int64Array `db:"tagIds"`
		CreatedAt   time.Time     `db:"createdAt,readonly"`
		PublishedAt time.Time     `db:"publishedAt"`
		StatusID    int           `db:"statusId"`
	}

	ShortNews struct {
		NewsID      int           `db:"newsId"`
		Title       string        `db:"title"`
		PublishedAt time.Time     `db:"publishedAt"`
		CategoryID  int           `db:"categoryId"`
		TagIds      pq.Int64Array `db:"tagIds"`
	}

	Categories struct {
		CategoryID  int    `db:"categoryId"`
		Title       string `db:"title"`
		OrderNumber int    `db:"orderNumber"`
		StatusID    int    `db:"statusId"`
	}

	Tags struct {
		TagID    int    `db:"tagId"`
		Title    string `db:"title"`
		StatusID int    `db:"statusId"`
	}

	statuses struct {
		StatusID int    `db:"statusId"`
		Title    string `db:"title"`
	}
)
