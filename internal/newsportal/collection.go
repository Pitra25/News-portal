package newsportal

import "News-portal/internal/db"

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

func (f *Filters) ToDB() db.Filters {
	return db.NewFilters(
		f.News.CategoryId, f.News.TagId,
		f.Page.PageSize, f.Page.Page,
	)
}

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

//go:generate colgen -imports=News-portal/internal/db
//colgen:News,Category,Tag
//colgen:News:UniqueTagIDs,MapP(db)
//colgen:Tag:MapP(db)
//colgen:Category:MapP(db)

// MapP converts slice of type T to slice of type M with given converter with pointers.
func MapP[T, M any](a []T, f func(*T) *M) []M {
	n := make([]M, len(a))
	for i := range a {
		n[i] = *f(&a[i])
	}
	return n
}

func (nl NewsList) SetTags(tags Tags) {
	tagIndex := tags.Index()
	for i, v := range nl {
		for _, tag := range v.TagIDs {
			if t, ok := tagIndex[tag]; ok {
				nl[i].Tags = append(nl[i].Tags, t)
			}
		}
	}
}

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
