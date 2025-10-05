package db

import (
	"github.com/go-pg/pg/v10"
)

type NewsRepo struct {
	db *pg.DB
}

func NewNews(db *pg.DB) *NewsRepo {
	return &NewsRepo{
		db: db,
	}
}

func (m *NewsRepo) GetByFilters(fil Filters) ([]News, error) {
	// formation of restrictions
	var (
		limit, offset = fil.Page.paginator()
		results       []News
	)

	if err := m.db.Model(&results).
		ColumnExpr(`DISTINCT ON ("tagIds") *`).
		Where(`"statusId" = ?`, newsStatus).
		Where(`"publishedAt" <= now()`).
		Where(`? = ANY("tagIds")`, fil.News.TagId).
		Where(`"categoryId" = ?`, fil.News.CategoryId).
		Limit(limit).Offset(offset).
		Select(); err != nil {
		return nil, err
	}

	return results, nil
}

func (m *NewsRepo) GetById(id int) (News, error) {

	result := News{NewsID: id}
	if err := m.db.Model(&result).
		Where(`"statusId" = ?`, newsStatus).
		Where(`"publishedAt" <= now()`).
		WherePK().
		Select(); err != nil {
		return result, err
	}

	return result, nil
}

func (m *NewsRepo) GetCount(filter Filters) (int, error) {

	count, err := m.db.Model((*News)(nil)).
		Where(`"statusId" = ?`, newsStatus).
		Where(`"categoryId" = ?`, filter.News.CategoryId).
		Where(`? = ANY("tagIds")`, filter.News.TagId).
		Where(`"publishedAt" <= now()`).
		Count()
	if err != nil {
		return count, err
	}

	return count, nil
}
