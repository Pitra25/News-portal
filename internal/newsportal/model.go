package newsportal

import (
	"News-portal/internal/db"
)

type (
	NewsFilters struct {
		CategoryId int
		TagId      int
	}

	PageFilters struct {
		PageSize int
		Page     int
	}

	Filters struct {
		News NewsFilters
		Page PageFilters
	}

	Tag struct{ db.Tag }

	Category struct{ db.Category }

	News struct {
		db.News
		Category *Category
		Tags     []Tag
	}
)

func NewFilters(categoryId, tagId, pageSize, page int) Filters {
	return Filters{
		NewsFilters{
			CategoryId: categoryId,
			TagId:      tagId,
		},
		PageFilters{
			PageSize: pageSize,
			Page:     page,
		},
	}
}
