package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
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

func filPubAt(orm *orm.Query) *orm.Query {
	return orm.Where(
		`"t".? <= now()`,
		pg.Ident(Columns.News.PublishedAt),
	)
}

func setQueryFilters(orm *orm.Query) *orm.Query {
	return orm.Where(
		`"t".? = ?`,
		pg.Ident(Columns.News.StatusID),
		newsStatus,
	)
}

func filters(orm *orm.Query, fil Filters) *orm.Query {
	query := setQueryFilters(filPubAt(orm))
	if fil.News.CategoryId != 0 {
		query.Where(
			`"t".? = ?`,
			pg.Ident(Columns.News.CategoryID),
			fil.News.CategoryId,
		)
	}
	if fil.News.TagId != 0 {
		query.Where(
			`? = ANY("t".?)`,
			fil.News.TagId,
			pg.Ident(Columns.News.TagIDs),
		)
	}
	return query
}
