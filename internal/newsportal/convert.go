package newsportal

import (
	"News-portal/internal/db"
)

func NewCategory(c db.Category) Category {
	return Category{
		CategoryID: c.ID,
		Title:      c.Title,
	}
}

func NewCategories(c []db.Category) []Category {
	var result []Category
	for _, v := range c {
		result = append(result, NewCategory(v))
	}
	return result
}

func NewNews(
	newsDB db.News,
	tags []Tag,
) News {
	return News{
		NewsID:      newsDB.ID,
		Title:       newsDB.Title,
		Content:     *newsDB.Content,
		Author:      newsDB.Author,
		Category:    NewCategory(*newsDB.Category),
		Tags:        tags,
		TagIds:      newsDB.TagIDs,
		PublishedAt: newsDB.PublishedAt,
	}
}

func NewTags(tagDB []db.Tag) []Tag {
	var tag []Tag
	for _, v := range tagDB {
		tag = append(tag, Tag{
			TagID: v.ID,
			Title: v.Title,
		})
	}
	return tag
}

func (f *Filters) ToDB() db.Filters {
	filter := db.NewFilters(
		f.News.NewsId, f.News.CategoryId, f.News.TagId,
		f.Page.PageSize, f.Page.Page,
	)
	return filter
}
