package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

const newsStatus = 1 // published

type Filters struct {
	CategoryId int
	TagId      int
	PageSize   int
	Page       int
}

func (fil *Filters) paginator() (int, int) {
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

func NewFilters(categoryId, tagId, pageSize, page int) Filters {
	return Filters{
		CategoryId: categoryId,
		TagId:      tagId,
		PageSize:   pageSize,
		Page:       page,
	}
}

func filPubAt(orm *orm.Query) *orm.Query {
	return orm.Where(
		`"t".? <= now()`,
		pg.Ident(Columns.News.PublishedAt),
	)
}

func filStatus(orm *orm.Query) *orm.Query {
	return orm.Where(
		`"t".? = ?`,
		pg.Ident(Columns.News.StatusID),
		newsStatus,
	)
}

func filters(orm *orm.Query, fil Filters) *orm.Query {
	query := filStatus(filPubAt(orm))
	if fil.CategoryId != 0 {
		query.Where(
			`"t".? = ?`,
			pg.Ident(Columns.News.CategoryID),
			fil.CategoryId,
		)
	}
	if fil.TagId != 0 {
		query.Where(
			`? = ANY("t".?)`,
			fil.TagId,
			pg.Ident(Columns.News.TagIDs),
		)
	}
	return query
}
