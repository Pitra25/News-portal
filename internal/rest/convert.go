package rest

import (
	"News-portal/internal/newsportal"
)

func NewCategories(c []newsportal.Category) []Category {
	if len(c) == 0 {
		return nil
	}

	var result []Category
	for _, v := range c {
		result = append(result, NewCategory(v))
	}
	return result
}

func NewCategory(c newsportal.Category) Category {
	if c.ID == 0 {
		return Category{}
	}

	return Category{
		ID:    c.ID,
		Title: c.Title,
	}
}

func NewTags(list []newsportal.Tag) []Tag {
	if len(list) == 0 {
		return nil
	}

	var tag []Tag
	for _, v := range list {
		tag = append(tag, NewTag(&v))
	}
	return tag
}

func NewTag(in *newsportal.Tag) Tag {
	if in == nil {
		return Tag{}
	}

	return Tag{
		ID:    in.ID,
		Title: in.Title,
	}
}

func NewNewsList(list []newsportal.News) []News {
	if len(list) == 0 {
		return nil
	}
	var news []News
	for _, v := range list {
		news = append(news, NewNews(&v))
	}
	return news
}

func NewNews(in *newsportal.News) News {
	if in == nil {
		return News{}
	}
	return News{
		ID:          in.ID,
		Title:       in.Title,
		Author:      in.Author,
		Content:     in.Content,
		PublishedAt: in.PublishedAt,
		Category:    NewCategory(in.Category),
		Tags:        NewTags(in.Tags),
	}
}

func NewNewsSummaries(list []newsportal.News) []NewsSummary {
	if len(list) == 0 {
		return nil
	}
	var news []NewsSummary
	for _, v := range list {
		news = append(news, NewNewsSummary(&v))
	}
	return news
}

func NewNewsSummary(in *newsportal.News) NewsSummary {
	if in == nil {
		return NewsSummary{}
	}
	return NewsSummary{
		ID:          in.ID,
		Title:       in.Title,
		PublishedAt: in.PublishedAt,
		Category:    NewCategory(in.Category),
		TagIds:      NewTags(in.Tags),
	}
}
