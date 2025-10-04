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

	newsStatus = 1 // published
)

type (
	NewsFilters struct {
		NewsId     int
		CategoryId int
		TagId      int
	}

	PageFilters struct {
		PageSize int
		Page     int
	}

	Filters struct {
		News NewsFilters
		Page PageFilters
	}

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
)

func (fil *NewsFilters) NewFilters() string {
	result := ` n."statusId" = :statusID AND n."publishedAt" <= now()`
	if fil.CategoryId != 0 {
		result = result + ` AND n."categoryId" = :categoryID`
	}
	if fil.TagId != 0 {
		result = result + ` AND :tagID = ANY(n."tagIds")`
	}
	return result
}

func (fil *PageFilters) paginator() (int, int) {
	limit, offset := 10, 0
	if fil.Page == 1 {
		limit = fil.PageSize
		offset = 0
	} else {
		limit = fil.PageSize
		offset = fil.Page * fil.PageSize
	}
	return limit, offset
}

func removeDuper(numbers pq.Int64Array) pq.Int64Array {
	seen := make(map[int64]bool)
	result := make([]int64, 0)

	for _, v := range numbers {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func NewFilters(
	newsId, categoryId, tagId,
	pageSize, page int,
) Filters {
	return Filters{
		NewsFilters{
			NewsId:     newsId,
			CategoryId: categoryId,
			TagId:      tagId,
		},
		PageFilters{
			PageSize: pageSize,
			Page:     page,
		},
	}
}
