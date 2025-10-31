package newsportal

import (
	"News-portal/internal/db"
	"time"
)

type (
	Filters struct {
		CategoryId int
		TagId      int
		TagIds     []int
		PageSize   int
		Page       int
	}

	UpdateStatus struct {
		NewsId     *int
		CategoryId *int
		TagId      *int
	}

	Tag      struct{ db.Tag }
	TagInput struct {
		ID    *int
		Title string
	}

	Category      struct{ db.Category }
	CategoryInput struct {
		ID          *int
		Title       string
		OrderNumber *int
	}

	News struct {
		db.News
		Category *Category
		Tags     []Tag
	}

	NewsInput struct {
		Id          *int
		Title       string
		Content     *string
		Author      string
		CategoryID  int
		TagIDs      []int
		PublishedAt *time.Time
	}
)

func (f *Filters) ToDB() *db.NewsSearch {
	//statusID := db.StatusEnabled
	//timeNow := time.Now()
	filter := db.NewsSearch{
		//StatusIDEQuals: &statusID,
		//PublishedAtLE:  &timeNow,
	}
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
	result := &db.News{
		Title:       in.Title,
		Content:     in.Content,
		Author:      in.Author,
		CategoryID:  in.CategoryID,
		TagIDs:      in.TagIDs,
		PublishedAt: *in.PublishedAt,
	}
	if in.Id != nil {
		result.ID = *in.Id
	}

	return result
}

func categoryToDB(in *CategoryInput) *db.Category {
	result := &db.Category{
		Title:       in.Title,
		OrderNumber: in.OrderNumber,
	}
	if in.ID != nil {
		result.ID = *in.ID
	}

	return result
}

func tagToDB(in *TagInput) *db.Tag {
	result := &db.Tag{
		Title: in.Title,
	}
	if in.ID != nil {
		result.ID = *in.ID
	}

	return result
}
