package db

import (
	"News-portal/output"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type NewsRepo struct {
	db *pg.DB
}

func NewNews(db *pg.DB) *NewsRepo {
	return &NewsRepo{
		db: db,
	}
}

func (m *NewsRepo) GetByFilters(fil Filters) ([]output.News, error) {
	// formation of restrictions
	var (
		limit, offset = fil.Page.paginator()
		results       []output.News
	)

	if err := m.db.Model(&results).
		Where(`"statusId" = ?`, newsStatus).
		Where(`"publishedAt" <= now()`).
		Where(`? = ANY("tagIds")`, fil.News.TagId).
		Where(`"categoryId" = ?`, fil.News.CategoryId).
		Limit(limit).Offset(offset).
		Select(); err != nil {
		return nil, err
	}

	for _, result := range results {
		removeDuper(&result)
	}

	return results, nil
}

func s(orm *orm.Query) *orm.Query {
	return orm.Where(`"statusId" = ?`, newsStatus).
		Where(`"publishedAt" <= now()`)
}

func (m *NewsRepo) GetById(id int) (output.News, error) {

	result := output.News{ID: id}
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

	count, err := m.db.Model((*output.News)(nil)).
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
