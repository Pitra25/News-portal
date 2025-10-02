package newsportal

import (
	"News-portal/internal/db"
)

func newsDtoToJson(
	newsDB db.News,
	tags []Tag,
	categoryDB db.Categories,
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

func shortNewsDtoToJson(
	newsDB db.ShortNews,
	tags []Tag,
	categoryDB db.Categories,
) ShortNews {
	return ShortNews{
		NewsID: newsDB.NewsID,
		Title:  newsDB.Title,
		Category: Category{
			CategoryID: categoryDB.CategoryID,
			Title:      categoryDB.Title,
		},
		TagIds:      tags,
		PublishedAt: newsDB.PublishedAt,
	}
}
