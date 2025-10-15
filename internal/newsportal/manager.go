package newsportal

import (
	"News-portal/internal/db"
	"fmt"

	"golang.org/x/net/context"
)

type Manager struct {
	repo db.NewsRepo
}

func NewManager(dbc *db.DB) *Manager {
	return &Manager{
		repo: db.NewNewsRepo(dbc),
	}
}

/*** News ***/

func (m *Manager) GetNewsByFilters(ctx context.Context, fil Filters) ([]News, error) {
	dbNews, err := m.repo.NewsByFilters(
		ctx, fil.filter(), *fil.pager(),
	)
	if err != nil {
		return nil, fmt.Errorf("news fetch failed: %w", err)
	}
	result := NewNewsList(dbNews)

	// collect everything in 1 news
	tags, err := m.GetTagsByID(ctx, result.UniqueTagIDs())
	if err != nil {
		return nil, fmt.Errorf("tags fetch failed: %w", err)
	}

	//collect everything in a news array
	result.SetTags(tags)

	return result, nil
}

func (m *Manager) GetNewsById(ctx context.Context, id int) (*News, error) {
	// receiving news by ID
	news, err := m.repo.NewsByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("news fetch failed: %w", err)
	}
	result := NewNews(news)

	tags, err := m.GetTagsByID(ctx, result.TagIDs)
	if err != nil {
		return nil, fmt.Errorf("tags fetch failed: %w", err)
	}

	result.Tags = tags

	return result, nil
}

func (m *Manager) GetNewsCount(ctx context.Context, fil Filters) (int, error) {
	return m.repo.CountNews(ctx, fil.filter())
}

/*** Category ***/

func (m *Manager) GetAllCategory(ctx context.Context) ([]Category, error) {
	categories, err := m.repo.CategoriesByFilters(ctx, &db.CategorySearch{}, db.PagerNoLimit)

	return NewCategories(categories), err
}

/*** Tag ***/

func (m *Manager) GetAllTag(ctx context.Context) ([]Tag, error) {
	tags, err := m.repo.TagsByFilters(ctx, &db.TagSearch{}, db.PagerNoLimit)

	return NewTags(tags), err
}

func (m *Manager) GetTagsByID(ctx context.Context, ids []int) (Tags, error) {
	fil := db.TagSearch{}
	if len(ids) > 0 {
		fil.IDs = ids
	}
	tags, err := m.repo.TagsByFilters(
		ctx, &fil, db.PagerNoLimit,
	)

	return NewTags(tags), err
}
