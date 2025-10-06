package newsportal

import (
	"News-portal/internal/db"
	"News-portal/output"
)

func NewCategory(c output.Category) Category {
	return Category{
		CategoryID: c.ID,
		Title:      c.Title,
	}
}

func NewCategoryArr(c []output.Category) []Category {
	var result []Category
	for _, v := range c {
		result = append(result, NewCategory(v))
	}
	return result
}

func NewNews(
	newsDB output.News,
	category Category,
	tags []Tag,
) News {
	return News{
		NewsID:      newsDB.ID,
		Title:       newsDB.Title,
		Content:     *newsDB.Content,
		Author:      newsDB.Author,
		Category:    category,
		Tags:        tags,
		TagIds:      newsDB.TagIDs,
		PublishedAt: newsDB.PublishedAt,
	}
}

func NewTagArr(tagDB []output.Tag) []Tag {
	var tag []Tag
	for _, v := range tagDB {
		tag = append(tag, Tag{
			TagID: v.ID,
			Title: v.Title,
		})
	}
	return tag
}

//func NewNewsArr(newsDB []db.News) []News {
//	var news []News
//	for _, v := range newsDB {
//		news = append(news, NewNews(v))
//	}
//	return news
//}

func NewShortNewsArr(
	newsDB output.News,
	categoryDB Category,
	tags []Tag,
) ShortNews {
	return ShortNews{
		NewsID:      newsDB.ID,
		Title:       newsDB.Title,
		PublishedAt: newsDB.PublishedAt,
		Category:    categoryDB,
		TagIds:      newsDB.TagIDs,
		Tags:        tags,
	}
}

func (f *Filters) ToDB() db.Filters {
	filter := db.NewFilters(
		f.News.NewsId, f.News.CategoryId, f.News.TagId,
		f.Page.PageSize, f.Page.Page,
	)
	return filter
}
