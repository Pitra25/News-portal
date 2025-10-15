package newsportal

import (
	"News-portal/internal/db"
	"time"
)

type (
	Filters struct {
		CategoryId int
		TagId      int
		PageSize   int
		Page       int
	}

	Tag      struct{ db.Tag }
	TagInput struct{ Title string }

	Category      struct{ db.Category }
	CategoryInput struct {
		Title       string
		OrderNumber *int
	}

	News struct {
		db.News
		Category *Category
		Tags     []Tag
	}

	NewsInput struct {
		Title       string
		Content     *string
		Author      string
		CategoryID  int
		TagIDs      []int
		PublishedAt *time.Time
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

func newsToDB(in *NewsInput) *db.News {
	return &db.News{
		Title:       in.Title,
		Content:     in.Content,
		Author:      in.Author,
		CategoryID:  in.CategoryID,
		TagIDs:      in.TagIDs,
		PublishedAt: *in.PublishedAt,
	}
}

func categoryToDB(in *CategoryInput) *db.Category {
	return &db.Category{
		Title:       in.Title,
		OrderNumber: in.OrderNumber,
	}
}

func tagToDB(in *TagInput) *db.Tag {
	return &db.Tag{
		Title: in.Title,
	}
}
