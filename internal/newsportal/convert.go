package newsportal

import "News-portal/internal/db"

func NewCategory(c db.Categories) Category {
	return Category{
		CategoryID: c.CategoryID,
		Title:      c.Title,
	}
}

func NewTag(tagDB db.Tags) Tag {
	return Tag{
		TagID: tagDB.TagID,
		Title: tagDB.Title,
	}
}

func NewNews(
	newsDB db.News,
	categoryDB Category,
	tags []Tag,
) News {
	return News{
		NewsID:  newsDB.NewsID,
		Title:   newsDB.Title,
		Content: newsDB.Content,
		Author:  newsDB.Author,
		Category: Category{
			CategoryID: categoryDB.CategoryID,
			Title:      categoryDB.Title,
		},
		Tags:        tags,
		TagIds:      newsDB.TagIds,
		PublishedAt: newsDB.PublishedAt,
	}
}

func NewFilterDB(fil Filters) db.Filters {
	filter := db.NewFilters(
		fil.News.NewsId, fil.News.CategoryId, fil.News.TagId,
		fil.Page.PageSize, fil.Page.Page,
	)
	return filter
}
