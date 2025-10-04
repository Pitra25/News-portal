package newsportal

import "time"

type (
	NewsFilters struct {
		NewsId     int
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

	Statuses struct {
		StatusID int
		Title    string
	}

	Tag struct {
		TagID int
		Title string
	}

	Category struct {
		CategoryID int
		Title      string
	}

	News struct {
		NewsID      int
		Title       string
		Content     string
		Author      string
		Category    Category
		TagIds      []int64
		Tags        []Tag
		PublishedAt time.Time
	}

	ShortNews struct {
		NewsID      int
		Title       string
		PublishedAt time.Time
		Category    Category
		Tags        []Tag
		TagIds      []int64
	}
)

func NewFilters(
	newsId, categoryId, tagId,
	pageSize, page int,
) Filters {
	return Filters{
		NewsFilters{
			NewsId:     newsId,
			CategoryId: categoryId,
			TagId:      tagId,
		},
		PageFilters{
			PageSize: pageSize,
			Page:     page,
		},
	}
}
