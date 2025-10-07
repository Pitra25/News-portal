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

	if err := filters(m.db.Model(&results), fil).
		Relation("Category").
		Limit(limit).Offset(offset).
		Select(); err != nil {
		return nil, err
	}

	for _, result := range results {
		removeDuper(&result)
	}
	return results, nil
}

func (m *NewsRepo) GetById(id int) (News, error) {

	result := News{ID: id}
	if err := filters(m.db.Model(&result), Filters{}).
		Relation("Category").
		WherePK().
		Select(); err != nil {
		return result, err
	}

	return result, nil
}

func (m *NewsRepo) GetCount(filter Filters) (int, error) {

	count, err := filters(m.db.Model((*News)(nil)), filter).
		Count()
	if err != nil {
		return count, err
	}

	return count, nil
}
