package newsportal

import (
	"News-portal/internal/db"
	"fmt"

	"golang.org/x/net/context"
)

type Manager struct {
	db *db.DB
}

func NewManager(db *db.DB) *Manager {
	return &Manager{
		db: db,
	}
}

func (m *Manager) GetNewsByFilters(ctx context.Context, fil Filters) ([]News, error) {

	dbNews, err := m.db.Repo.GetNewsByFilters(ctx, fil.ToDB())
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
	news, err := m.db.Repo.GetNewsById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("news fetch failed: %w", err)
	}

	tags, err := m.GetTagsByID(ctx, news.TagIDs)
	if err != nil {
		return nil, fmt.Errorf("tags fetch failed: %w", err)
	}

	result := NewNews(news)
	if result == nil {
		return nil, nil
	}
	result.Tags = tags

	return result, nil
}

func (m *Manager) GetTagsByID(ctx context.Context, ids []int) (Tags, error) {
	tags, err := m.db.Repo.GetTagByID(ctx, ids)
	return NewTags(tags), err
}

func (m *Manager) GetNewsCount(ctx context.Context, fil Filters) (int, error) {
	return m.db.Repo.GetNewsCount(ctx, fil.ToDB())
}

func (m *Manager) GetAllTag(ctx context.Context) ([]Tag, error) {
	tags, err := m.db.Repo.GetTagsList(ctx)
	return NewTags(tags), err
}

func (m *Manager) GetAllCategory(ctx context.Context) ([]Category, error) {
	categories, err := m.db.Repo.GetListCategories(ctx)
	return NewCategories(categories), err
}
