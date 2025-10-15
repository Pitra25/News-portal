package newsportal

import "News-portal/internal/db"

type (
	Filters struct {
		CategoryId int
		TagId      int
		PageSize   int
		Page       int
	}

	Tag struct{ db.Tag }

	Category struct{ db.Category }

	News struct {
		db.News
		Category *Category
		Tags     []Tag
	}
)

func (f *Filters) filter() *db.NewsSearch {
	filter := db.NewsSearch{}
	if f.TagId != 0 {
		filter.IDs = []int{f.TagId}
	}
	if f.CategoryId != 0 {
		filter.CategoryID = &f.CategoryId
	}
	return &filter
}

func (f *Filters) pager() *db.Pager {
	pager := db.Pager{}
	if f.Page != 0 {
		pager.Page = f.Page
	}
	if f.PageSize != 0 {
		pager.PageSize = f.PageSize
	}
	return &pager
}

func NewFilters(categoryId, tagId, pageSize, page int) Filters {
	return Filters{
		CategoryId: categoryId,
		TagId:      tagId,
		PageSize:   pageSize,
		Page:       page,
	}
}
