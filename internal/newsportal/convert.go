package newsportal

import (
	"News-portal/internal/db"
)

func NewCategory(in db.Category) *Category {
	return &Category{
		Category: in,
	}
}

func NewCategories(list []db.Category) []Category {
	if list == nil {
		return nil
	}
	var result []Category
	for _, v := range list {
		result = append(result, *NewCategory(v))
	}
	return result
}

func NewNewsList(list []db.News) NewsList {
	if list == nil {
		return nil
	}
	var results []News
	for _, item := range list {
		n := NewNews(&item)
		if n == nil {
			break
		}
		results = append(results, *n)
	}
	return results
}

func NewNews(in *db.News) *News {
	if in.Category == nil || in.ID == 0 {
		return nil
	}
	return &News{
		News:     *in,
		Category: *NewCategory(*in.Category),
	}
}

func NewTags(list []db.Tag) Tags {
	if len(list) == 0 {
		return nil
	}
	var tag []Tag
	for _, v := range list {
		tag = append(tag, NewTag(v))
	}
	return tag
}

func NewTag(in db.Tag) Tag {
	if in.ID == 0 {
		return Tag{}
	}
	return Tag{Tag: in}
}

func (f *Filters) ToDB() db.Filters {
	return db.NewFilters(
		f.News.NewsId, f.News.CategoryId, f.News.TagId,
		f.Page.PageSize, f.Page.Page,
	)
}
