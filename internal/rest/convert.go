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

func NewCategoryArr(c []newsportal.Category) []Category {
	var result []Category
	for _, v := range c {
		result = append(result, NewCategory(v))
	}
	return result
}

func NewTag(tagDB []newsportal.Tag) []Tag {
	var tag []Tag
	for _, v := range tagDB {
		tag = append(tag, Tag{
			TagID: v.TagID,
			Title: v.Title,
		})
	}
	return tag
}

func NewNewsArr(newsDB []newsportal.News) []News {
	var news []News
	for _, v := range newsDB {
		news = append(news, NewNews(v))
	}
	return news
}

func NewShortNews(newsDB []newsportal.News) []ShortNews {
	var news []ShortNews
	for _, v := range newsDB {
		news = append(news, ShortNews{
			NewsID:      v.NewsID,
			Title:       v.Title,
			PublishedAt: v.PublishedAt,
			Category:    NewCategory(v.Category),
			TagIds:      NewTag(v.Tags),
		})
	}
	return news
}

func NewNews(v newsportal.News) News {
	return News{
		NewsID:      v.NewsID,
		Title:       v.Title,
		Content:     v.Content,
		Author:      v.Author,
		Category:    NewCategory(v.Category),
		Tags:        NewTag(v.Tags),
		PublishedAt: v.PublishedAt,
	}
}
