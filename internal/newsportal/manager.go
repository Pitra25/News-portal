package newsportal

import (
	"News-portal/internal/db"
)

type Manager struct {
	db *db.DB
}

func NewManager(db *db.DB) *Manager {
	return &Manager{
		db: db,
	}
}

func (m *Manager) GetNewsByFilters(fil Filters) ([]News, error) {

	filter := fil.ToDB()
	newsDB, err := m.db.News.GetByFilters(filter)
	if err != nil {
		return nil, err
	}

	// collect everything in 1 news
	tagsArr, _, _, err := getTagsAndCategory(m, newsDB)
	if err != nil {
		return nil, err
	}

	// collect everything in a news array
	var result []News
	for _, v := range newsDB {
		// find an object with the necessary tags
		tags := GetNewsMetadata(tagsArr, v)

		// Collect everything in 1 news
		result = append(result, NewNews(v, tags))
	}

	return result, nil
}

func (m *Manager) GetNewsById(id int) (News, error) {
	// receiving news by ID
	data, err := m.db.News.GetById(id)
	if err != nil {
		return News{}, err
	}

	// Get the name of the news tags
	var tagIds []int
	for _, v := range data.TagIDs {
		tagIds = append(tagIds, int(v))
	}
	tags, err := m.db.Tags.GetByID(tagIds)
	if err != nil {
		return News{}, err
	}

	// collect everything in 1 news item
	return NewNews(
		data,
		NewTags(tags),
	), nil

}

func (m *Manager) GetNewsCount(fil Filters) (int, error) {
	filter := fil.ToDB()
	return m.db.News.GetCount(filter)
}

func (m *Manager) GetAllTag() ([]Tag, error) {
	tags, err := m.db.Tags.GetAll()
	if err != nil {
		return nil, err
	}

	return NewTags(tags), nil
}

func (m *Manager) GetAllCategory() ([]Category, error) {
	data, err := m.db.Categories.GetAll()
	if err != nil {
		return nil, err
	}

	var result []Category
	for _, v := range data {
		result = append(result, NewCategory(v))
	}

	return result, nil
}
