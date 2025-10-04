package rest

import (
	"News-portal/internal/newsportal"
)

func NewCategory(c newsportal.Category) Category {
	return Category{
		CategoryID: c.CategoryID,
		Title:      c.Title,
	}
}

func NewTag(tagDB newsportal.Tag) Tag {
	return Tag{
		TagID: tagDB.TagID,
		Title: tagDB.Title,
	}
}

func NewNews(
	newsDB newsportal.News,
	tags []Tag,
) News {
	return News{
		NewsID:  newsDB.NewsID,
		Title:   newsDB.Title,
		Content: newsDB.Content,
		Author:  newsDB.Author,
		Category: Category{
			CategoryID: newsDB.Category.CategoryID,
			Title:      newsDB.Category.Title,
		},
		Tags:        tags,
		PublishedAt: newsDB.PublishedAt,
	}
}

func NewShortNews(
	newsDB newsportal.ShortNews,
	tags []Tag,
) ShortNews {
	return ShortNews{
		NewsID:      newsDB.NewsID,
		Title:       newsDB.Title,
		PublishedAt: newsDB.PublishedAt,
		Category: Category{
			CategoryID: newsDB.Category.CategoryID,
			Title:      newsDB.Category.Title,
		},
		TagIds: tags,
	}
}
