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
		tableName  struct{} `pg:"newsportal.news"`
		NewsID     int      `pg:"newsId,pk"`
		Title      string
		Content    string
		Author     string
		CategoryID int `pg:"categoryId"`
		//Category    *Categories   `pg:"rel:has-one,fk:categoryId"`
		TagIds      pq.Int64Array `pg:"tagIds"`
		CreatedAt   time.Time     `pg:"createdAt"`
		PublishedAt time.Time     `pg:"publishedAt"`
		StatusID    int           `pg:"statusId"`
		//Status      *Statuses     `pg:"rel:has-one,fk:statusId"`
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
		//Status      *Statuses `pg:"rel:has-one,fk:statusId"`
	}

	Tags struct {
		tableName struct{} `pg:"newsportal.tags"`
		TagID     int      `pg:"tagId,pk"`
		Title     string
		StatusID  int `pg:"statusId"`
		//Status   *Statuses `pg:"rel:has-one,fk:statusId"`
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
