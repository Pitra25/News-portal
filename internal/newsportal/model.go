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

func NewFilters(categoryId, tagId, pageSize, page int) Filters {
	return Filters{
		CategoryId: categoryId,
		TagId:      tagId,
		PageSize:   pageSize,
		Page:       page,
	}
}
