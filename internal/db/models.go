package db

import (
	"News-portal/output"
)

const newsStatus = 1 // published

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
)

func (fil *PageFilters) paginator() (int, int) {
	limit, offset := 10, 0
	if fil.Page == 1 {
		limit = fil.PageSize
		offset = 0
	} else {
		limit = fil.PageSize
		offset = fil.Page * fil.PageSize
	}
	return limit, offset
}

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

func removeDuper(news *output.News) *output.News {
	seen := make(map[int]bool)
	result := make([]int, 0)

	for _, v := range news.TagIDs {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	news.TagIDs = result

	return news
}
