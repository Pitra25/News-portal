package newsportal

import (
	"News-portal/internal/db"
)

func NewCategory(in *db.Category) *Category {
	if in == nil {
		return nil
	}

	return &Category{
		Category: *in,
	}
}

func NewNews(in *db.News) *News {
	if in == nil {
		return nil
	}

	return &News{
		News:     *in,
		Category: NewCategory(in.Category),
	}
}

func NewTag(in *db.Tag) *Tag {
	if in == nil {
		return nil
	}

	return &Tag{Tag: *in}
}

func (f *Filters) ToDB() db.Filters {
	return db.NewFilters(
		f.News.NewsId, f.News.CategoryId, f.News.TagId,
		f.Page.PageSize, f.Page.Page,
	)
}
