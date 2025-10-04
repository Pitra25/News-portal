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

	// формирование массивов id категорий и тегов новости
	tagIds := map[int][]int64{}
	tagIdsArr := []int{}
	categoryIdArr := []int{}
	for _, v := range newsDB {
		tagIds[v.NewsID] = v.TagIds
		tagIdsArr = append(tagIdsArr, v.NewsID)
		categoryIdArr = append(categoryIdArr, v.CategoryID)
	}

	// Get the name of the news tags
	tags, err := m.db.Tags.GetByID(tagIdsArr)
	if err != nil {
		return []News{}, err
	}
	// transform news into a new type
	newTag := []Tag{}
	for _, v := range tags {
		newTag = append(newTag, NewTag(v))
	}

	// getting the name of the news category
	newCategory, err := m.db.Categories.GetById(categoryIdArr)
	if err != nil {
		return []News{}, err
	}

	// собрать все в массив новостей
	result := []News{}
	// обойти массив новостей
	for _, v := range newsDB {
		// найти объект с нужными тегами
		tags := []Tag{}
		for _, tag := range newTag {
			if tag.TagID == int(v.TagIds[0]) {
				tags = append(tags, tag)
			}
		}

		// найти объект с нужной категорией
		category := Category{}
		for _, cat := range newCategory {
			if cat.CategoryID == v.CategoryID {
				category = NewCategory(cat)
			}
		}

		// собрать все в 1 новости
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

	// формирование массивов id категорий и тегов новости
	tagIds := map[int][]int64{}
	tagIdsArr := []int{}
	categoryIdArr := []int{}
	for _, v := range newsDB {
		tagIds[v.NewsID] = v.TagIds
		tagIdsArr = append(tagIdsArr, v.NewsID)
		categoryIdArr = append(categoryIdArr, v.CategoryID)
	}

	// Get the name of the news tags
	tags, err := m.db.Tags.GetByID(tagIdsArr)
	if err != nil {
		return []ShortNews{}, err
	}
	// transform news into a new type
	newTag := []Tag{}
	for _, v := range tags {
		newTag = append(newTag, NewTag(v))
	}

	// getting the name of the news category
	newCategory, err := m.db.Categories.GetById(categoryIdArr)
	if err != nil {
		return []ShortNews{}, err
	}

	// собрать все в массив новостей
	result := []ShortNews{}
	// обойти массив новостей
	for _, v := range newsDB {
		// найти объект с нужными тегами
		tags := []Tag{}
		for _, tag := range newTag {
			if tag.TagID == int(v.TagIds[0]) {
				tags = append(tags, tag)
			}
		}

		// найти объект с нужной категорией
		category := Category{}
		for _, cat := range newCategory {
			if cat.CategoryID == v.CategoryID {
				category = NewCategory(cat)
			}
		}

		// собрать все в 1 новости
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

	tagIds := []int{}
	for _, v := range data.TagIds {
		tagIds = append(tagIds, int(v))
	}
	// Get the name of the news tags
	tags, err := m.db.Tags.GetByID(tagIds)
	if err != nil {
		return News{}, err
	}
	// transform news into a new type
	newTag := []Tag{}
	for _, v := range tags {
		newTag = append(newTag, NewTag(v))
	}

	categoryId := []int{}
	categoryId = append(categoryId, data.CategoryID)
	// getting the name of the news category
	category, err := m.db.Categories.GetById(categoryId)
	if err != nil {
		return News{}, err
	}

	// collect everything in 1 news item
	return NewNews(
		data,
		newTag,
		NewCategory(category[0]),
	), nil

}

func (m *Manager) GetNewsCount(fil Filters) (int, error) {
	filter := NewFilterDB(fil)
	return m.db.News.GetCount(filter)
}

func (m *Manager) GetAllTag() ([]Tag, error) {
	tags, err := m.db.Tags.GetAll()
	if err != nil {
		return nil, err
	}

	var result = make([]Tag, len(tags))
	for _, tag := range tags {
		result = append(result, NewTag(tag))
	}

	return result, nil
}

func (m *Manager) GetAllCategory() ([]Category, error) {
	data, err := m.db.Categories.GetAll()
	if err != nil {
		return nil, err
	}

	var result = make([]Category, len(data))
	for _, v := range data {
		result = append(result, NewCategory(v))
	}

	return result, nil
}
