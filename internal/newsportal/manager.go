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
		ctx, fil.ToDB(), db.Pager{Page: fil.Page, PageSize: fil.PageSize},
		db.WithColumns(db.Columns.News.Category),
	)
	if err != nil {
		return nil, fmt.Errorf("news fetch failed: %w", err)
	}
	result := NewNewsList(dbNews)

	// collect everything in 1 news
	tags, err := m.GetTagsByFilters(ctx, Filters{TagIds: result.UniqueTagIDs()})
	if err != nil {
		return nil, fmt.Errorf("tags fetch failed: %w", err)
	}

	//collect everything in a news array
	result.SetTags(tags)

	return result, nil
}

func (m *Manager) GetNewsByID(ctx context.Context, id int) (*News, error) {
	// receiving news by ID
	news, err := m.repo.NewsByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("news fetch failed: %w", err)
	}
	result := NewNews(news)

	tags, err := m.GetTagsByFilters(ctx, Filters{TagIds: result.TagIDs})
	if err != nil {
		return nil, fmt.Errorf("tags fetch failed: %w", err)
	}

	result.Tags = tags

	return result, nil
}

func (m *Manager) GetNewsCount(ctx context.Context, fil Filters) (int, error) {
	return m.repo.CountNews(ctx, fil.ToDB())
}

func (m *Manager) AddNews(ctx context.Context, in *NewsInput) (*News, error) {
	res, err := m.repo.AddNews(ctx, newsToDB(in))
	return NewNews(res), err
}

func (m *Manager) UpdateNews(ctx context.Context, in *NewsInput) (bool, error) {
	return m.repo.UpdateNews(ctx, newsToDB(in))
}

func (m *Manager) DeleteNews(ctx context.Context, id int) (bool, error) {
	return m.repo.DeleteNews(ctx, id)
}

/*** Category ***/

func (m *Manager) GetAllCategory(ctx context.Context) ([]Category, error) {
	categories, err := m.repo.CategoriesByFilters(ctx, &db.CategorySearch{}, db.PagerNoLimit)

	return NewCategories(categories), err
}

func (m *Manager) AddCategory(ctx context.Context, in *CategoryInput) (*Category, error) {
	if in.OrderNumber == nil {
		on, err := m.repo.MaxOrderNumber(ctx)
		if err != nil {
			return nil, err
		}
		on++
		in.OrderNumber = &on
	}

	res, err := m.repo.AddCategory(ctx, categoryToDB(in))
	return NewCategory(res), err
}

func (m *Manager) UpdateCategory(ctx context.Context, in *CategoryInput) (bool, error) {
	return m.repo.UpdateCategory(ctx, categoryToDB(in))
}

func (m *Manager) DeleteCategory(ctx context.Context, id int) (bool, error) {
	return m.repo.DeleteCategory(ctx, id)
}

/*** Tag ***/

func (m *Manager) GetTagsByFilters(ctx context.Context, fil Filters) (Tags, error) {
	search := db.TagSearch{}
	if len(fil.TagIds) > 0 {
		search.IDs = fil.TagIds
	}
	tags, err := m.repo.TagsByFilters(
		ctx, &search, db.PagerNoLimit,
	)

	return NewTags(tags), err
}

func (m *Manager) AddTag(ctx context.Context, in *TagInput) (*Tag, error) {
	res, err := m.repo.AddTag(ctx, tagToDB(in))
	return NewTag(res), err
}

func (m *Manager) UpdateTag(ctx context.Context, in *TagInput) (bool, error) {
	return m.repo.UpdateTag(ctx, tagToDB(in))
}

func (m *Manager) DeleteTag(ctx context.Context, id int) (bool, error) {
	return m.repo.DeleteTag(ctx, id)
}
