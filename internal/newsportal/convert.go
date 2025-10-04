package newsportal

import "News-portal/internal/db"

func NewCategory(c db.Categories) Category {
	return Category{
		CategoryID: c.CategoryID,
		Title:      c.Title,
	}
}

func NewNews(
	newsDB db.News,
	tags []Tag,
	categoryDB Category,
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

func NewNewsShort(
	newsDB db.News,
	tags []Tag,
	tagIds []int64,
	categoryDB db.Categories,
) ShortNews {
	return ShortNews{
		NewsID: newsDB.NewsID,
		Title:  newsDB.Title,
		Category: Category{
			CategoryID: categoryDB.CategoryID,
			Title:      categoryDB.Title,
		},
		TagIds:      tagIds,
		Tags:        tags,
		PublishedAt: newsDB.PublishedAt,
	}
}

func NewTag(tagDB db.Tags) Tag {
	return Tag{
		TagID: tagDB.TagID,
		Title: tagDB.Title,
	}
}

func NewFilterDB(fil Filters) db.Filters {
	filter := db.NewFilters(
		fil.News.NewsId, fil.News.CategoryId, fil.News.TagId,
		fil.Page.PageSize, fil.Page.Page,
	)
	return filter
}
