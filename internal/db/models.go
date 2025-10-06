package db

import (
	"time"

	"github.com/lib/pq"
)

const newsStatus = 1 // published

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
		tableName   struct{} `pg:"newsportal.news"`
		NewsID      int      `pg:"newsId,pk"`
		Title       string
		Content     string
		Author      string
		CategoryID  int           `pg:"categoryId"`
		TagIds      pq.Int64Array `pg:"tagIds"`
		CreatedAt   time.Time     `pg:"createdAt"`
		PublishedAt time.Time     `pg:"publishedAt"`
		StatusID    int           `pg:"statusId"`
	}

	Statuses struct {
		tableName struct{} `pg:"newsportal.statuses"`
		StatusID  int      `pg:"statusId,pk"`
		Title     string
	}

	Categories struct {
		tableName   struct{} `pg:"newsportal.categories"`
		CategoryID  int      `pg:"categoryId,pk"`
		Title       string
		OrderNumber int `pg:"orderNumber"`
		StatusID    int `pg:"statusId"`
	}

	Tags struct {
		tableName struct{} `pg:"newsportal.tags"`
		TagID     int      `pg:"tagId,pk"`
		Title     string
		StatusID  int `pg:"statusId"`
	}
)

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
