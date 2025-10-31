package rest

import (
	"News-portal/internal/newsportal"
)

//go:generate colgen -imports=News-portal/internal/newsportal -funcpkg=newsportal
//colgen:News,Category,Tag
//colgen:News:MapP(newsportal)
//colgen:Tag:MapP(newsportal)
//colgen:Category:MapP(newsportal)

func NewCategory(in *newsportal.Category) *Category {
	if in == nil {
		return nil
	}
	return &Category{
		ID:    in.ID,
		Title: in.Title,
	}
}

func NewNews(in *newsportal.News) *News {
	if in == nil {
		return nil
	}
	return &News{
		ID:          in.ID,
		Title:       in.Title,
		Author:      in.Author,
		Content:     in.Content,
		PublishedAt: in.PublishedAt,
		Category:    NewCategory(in.Category),
		Tags:        NewTags(in.Tags),
	}
}

func NewTag(in *newsportal.Tag) *Tag {
	if in == nil {
		return nil
	}
	return &Tag{
		ID:    in.ID,
		Title: in.Title,
	}
}

func NewNewsSummaries(in []newsportal.News) []NewsSummary {
	if in == nil {
		return nil
	}
	var result = make([]NewsSummary, len(in), len(in))
	for _, i := range in {
		result = append(result, *NewNewsSummary(&i))
	}
	return result
}

func NewNewsSummary(in *newsportal.News) *NewsSummary {
	if in == nil {
		return nil
	}
	return &NewsSummary{
		ID:          in.ID,
		Title:       in.Title,
		PublishedAt: in.PublishedAt,
		Category:    NewCategory(in.Category),
		TagIds:      NewTags(in.Tags),
	}
}
