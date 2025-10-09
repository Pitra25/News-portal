package db

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type Repo struct {
	db *pg.DB
}

func NewRepo(db *pg.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (m *Repo) GetNewsByFilters(ctx context.Context, fil Filters) ([]News, error) {
	// formation of restrictions
	var (
		limit, offset = fil.Page.paginator()
		results       []News
	)

	q := m.db.ModelContext(ctx, &results)
	if err := filters(q, fil).
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

func (m *Repo) GetNewsById(ctx context.Context, id int) (News, error) {

	result := News{ID: id}

	q := m.db.ModelContext(ctx, &result)
	err := filters(q, Filters{}).
		Relation("Category").
		WherePK().
		Select()

	return result, err
}

func (m *Repo) GetNewsCount(ctx context.Context, filter Filters) (int, error) {

	q := m.db.ModelContext(ctx, (*News)(nil))
	count, err := filters(q, filter).Count()

	return count, err
}

func (m *Repo) GetListCategories(ctx context.Context) ([]Category, error) {
	var list []Category

	q := m.db.ModelContext(ctx, &list)
	err := setQueryFilters(q).Select()

	return list, err
}

func (m *Repo) GetTagsList(ctx context.Context) ([]Tag, error) {
	var list []Tag

	q := m.db.ModelContext(ctx, &list)
	err := setQueryFilters(q).Select()

	return list, err
}

func (m *Repo) GetTagByID(ctx context.Context, ids []int) ([]Tag, error) {

	if len(ids) == 0 {
		return nil, nil
	}

	var list []Tag
	if err := setQueryFilters(m.db.ModelContext(ctx, &list)).
		Where(`"t"."tagId" IN (?)`, pg.In(ids)).
		Select(); err != nil {
		return nil, err
	}

	return list, nil
}
