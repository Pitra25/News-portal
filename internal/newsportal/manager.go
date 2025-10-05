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

	filter := NewFilterDB(fil)
	newsDB, err := m.db.News.GetByFilters(filter)
	if err != nil {
		return nil, err
	}

	// collect everything in 1 news
	tagsArr, categoryArr, tagIds, err := getTagsAndCategory(m, newsDB)
	if err != nil {
		return nil, err
	}

	// collect everything in a news array
	var result []News
	for _, v := range newsDB {
		// find an object with the necessary tags
		tags, category := GetNewsMetadata(tagsArr, categoryArr, v)

		// Collect everything in 1 news
		result = append(result, News{
			NewsID:      v.NewsID,
			Title:       v.Title,
			Content:     v.Content,
			Author:      v.Author,
			Category:    category,
			Tags:        tags,
			TagIds:      tagIds[v.NewsID],
			PublishedAt: v.PublishedAt,
		})
	}

	return result, nil
}

func (m *Manager) GetALlShortNewsByFilters(fil Filters) ([]ShortNews, error) {

	filter := NewFilterDB(fil)
	newsDB, err := m.db.News.GetByFilters(filter)
	if err != nil {
		return nil, err
	}

	// creating arrays of id categories and news tags
	tagsArr, categoryArr, tagIds, err := getTagsAndCategory(m, newsDB)
	if err != nil {
		return nil, err
	}

	// collect everything in a news array
	var result []ShortNews
	for _, v := range newsDB {
		// find an object with the necessary tags
		tags := []Tag{}
		for _, tag := range tagsArr {
			if tag.TagID == int(v.TagIds[0]) {
				tags = append(tags, tag)
			}
		}

		// find an object with the desired category
		category := Category{}
		for _, cat := range categoryArr {
			if cat.CategoryID == v.CategoryID {
				category = NewCategory(cat)
			}
		}

		// collect everything in 1 news item
		result = append(result, ShortNews{
			NewsID:      v.NewsID,
			Title:       v.Title,
			Category:    category,
			Tags:        tags,
			TagIds:      tagIds[v.NewsID],
			PublishedAt: v.PublishedAt,
		})
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
	var tagIds []int64
	for _, v := range data.TagIds {
		tagIds = append(tagIds, v)
	}
	tags, err := m.db.Tags.GetByID(tagIds)
	if err != nil {
		return News{}, err
	}

	// transform news into a new type
	var newTag []Tag
	for _, v := range tags {
		newTag = append(newTag, NewTag(v))
	}

	// getting the name of the news category
	category, err := m.db.Categories.GetById([]int{data.CategoryID})
	if err != nil {
		return News{}, err
	}

	// collect everything in 1 news item
	return NewNews(
		data,
		NewCategory(category[0]),
		newTag,
	), nil

}

func (m *Manager) GetNewsCount(fil Filters) (int, error) {
	filter := NewFilterDB(fil)
	return m.db.News.GetCount(filter)
}

func (m *Manager) GetAllTag() ([]Tag, error) {
	data, err := m.db.Tags.GetAll()
	if err != nil {
		return nil, err
	}

	var result []Tag
	for _, tag := range data {
		result = append(result, NewTag(tag))
	}

	return result, nil
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
