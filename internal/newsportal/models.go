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
		TagIds      []int
		Tags        []Tag
		PublishedAt time.Time
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

func getTagsAndCategory(m *Manager, data []db.News) ([]Tag, []Category, map[int][]int, error) {

	// Collecting IDs for requests
	var (
		tagIdsArr     []int
		categoryIdArr []int
	)

	tagIds := map[int][]int{}
	for _, v := range data {
		tagIds[v.ID] = v.TagIDs
		categoryIdArr = append(categoryIdArr, v.CategoryID)
		for _, tagId := range v.TagIDs {
			tagIdsArr = append(tagIdsArr, int(tagId))
		}
	}

	// Getting the tags
	var (
		tag        []Tag
		categories []Category
	)
	tagData, err := m.db.Tags.GetByID(tagIdsArr)
	if err != nil {
		return tag, categories, tagIds, err
	}

	// We get the categories
	categoriesDB, err := m.db.Categories.GetById(categoryIdArr)
	if err != nil {
		return tag, categories, tagIds, err
	}

	return NewTags(tagData), NewCategories(categoriesDB), tagIds, nil
}

func GetNewsMetadata(tagsArr []Tag, v db.News) []Tag {
	// find an object with the necessary tags
	var tags []Tag
	for _, tag := range tagsArr {
		for _, v := range v.TagIDs {
			if tag.TagID == v {
				tags = append(tags, tag)
			}
		}
	}

	return tags
}
