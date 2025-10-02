package newsportal

import "time"

type (
	Statuses struct {
		StatusID int
		Title    string
	}

	Tag struct {
		TagID int
		Title string
	}

	Category struct {
		CategoryID int
		Title      string
	}

	News struct {
		NewsID      int
		Title       string
		Content     string
		Author      string
		Category    Category
		TagIds      []int64
		Tags        []Tag
		PublishedAt time.Time
	}

	ShortNews struct {
		NewsID      int
		Title       string
		PublishedAt time.Time
		Category    Category
		TagIds      []Tag
	}

	Categories struct {
		CategoryID  int
		Title       string
		OrderNumber int
	}
)
