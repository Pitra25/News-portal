package rest

import "time"

type (
	Statuses struct {
		StatusID int    `json:"statusId"`
		Title    string `json:"title"`
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

	Categories struct {
		CategoryID  int    `json:"categoryId"`
		Title       string `json:"title"`
		OrderNumber int    `json:"orderNumber"`
	}
)
