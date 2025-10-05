package newsportal

import (
	"News-portal/internal/db"
	"time"
)

type (
	NewsFilters struct {
		NewsId     int
		CategoryId int
		TagId      int
	}

	PageFilters struct {
		PageSize int
		Page     int
	}

	Filters struct {
		News NewsFilters
		Page PageFilters
	}

	Statuses struct {
		StatusID int
		Title    string
	}

	Tag struct {
		TagID int
		Title string
	}

	Category struct {
		CategoryID int
		Title      string
	}

	News struct {
		NewsID      int
		Title       string
		Content     string
		Author      string
		Category    Category
		TagIds      []int64
		Tags        []Tag
		PublishedAt time.Time
	}

	ShortNews struct {
		NewsID      int
		Title       string
		PublishedAt time.Time
		Category    Category
		Tags        []Tag
		TagIds      []int64
	}
)

func NewFilters(
	newsId, categoryId, tagId,
	pageSize, page int,
) Filters {
	return Filters{
		NewsFilters{
			NewsId:     newsId,
			CategoryId: categoryId,
			TagId:      tagId,
		},
		PageFilters{
			PageSize: pageSize,
			Page:     page,
		},
	}
}

func getTagsAndCategory(m *Manager, data []db.News) ([]Tag, []db.Categories, map[int][]int64, error) {

	// Collecting IDs for requests
	var (
		tagIdsArr     []int64
		categoryIdArr []int
	)

	tagIds := map[int][]int64{}
	for _, v := range data {
		tagIds[v.NewsID] = v.TagIds
		categoryIdArr = append(categoryIdArr, v.CategoryID)
		for _, tagId := range v.TagIds {
			tagIdsArr = append(tagIdsArr, tagId)
		}
	}

	// Getting the tags
	var (
		tag        []Tag
		categories []db.Categories
	)
	tagData, err := m.db.Tags.GetByID(tagIdsArr)
	if err != nil {
		return tag, categories, tagIds, err
	}

	// Converting tags
	for _, v := range tagData {
		tag = append(tag, NewTag(v))
	}

	// We get the categories
	categories, err = m.db.Categories.GetById(categoryIdArr)
	if err != nil {
		return tag, categories, tagIds, err
	}

	return tag, categories, tagIds, nil
}

func GetNewsMetadata(tagsArr []Tag, categoryArr []db.Categories, v db.News) ([]Tag, Category) {
	// find an object with the necessary tags
	var tags []Tag
	for i, tag := range tagsArr {
		if tag.TagID == int(v.TagIds[i]) {
			tags = append(tags, tag)
		}
	}

	// find an object with the desired category
	var category Category
	for _, cat := range categoryArr {
		if cat.CategoryID == v.CategoryID {
			category = NewCategory(cat)
			break
		}
	}

	return tags, category
}
