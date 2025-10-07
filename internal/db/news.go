package db

import (
	"News-portal/output"

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

func (m *NewsRepo) GetByFilters(fil Filters) ([]output.News, error) {
	// formation of restrictions
	var (
		limit, offset = fil.Page.paginator()
		results       []output.News
	)

	if err := filters(m.db.Model(&results)).
		Relation("Category.title").
		Where(`? = ANY("t"."tagIds")`, fil.News.TagId).
		Where(`"t"."categoryId" = ?`, fil.News.CategoryId).
		Limit(limit).Offset(offset).
		Select(); err != nil {
		return nil, err
	}

	for _, result := range results {
		removeDuper(&result)
	}
	return results, nil
}

func (m *NewsRepo) GetById(id int) (output.News, error) {

	result := output.News{ID: id}
	if err := filters(m.db.Model(&result)).
		WherePK().
		Select(); err != nil {
		return result, err
	}

	return result, nil
}

func (m *NewsRepo) GetCount(filter Filters) (int, error) {

	count, err := filters(m.db.Model((*output.News)(nil))).
		Where(`"t"."categoryId" = ?`, filter.News.CategoryId).
		Where(`? = ANY("t"."tagIds")`, filter.News.TagId).
		Count()
	if err != nil {
		return count, err
	}

	return count, nil
}
