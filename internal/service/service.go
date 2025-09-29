package service

import (
	"News-portal/internal/db"
	"News-portal/internal/db/models"
	"News-portal/internal/service/methods"
)

type News interface {
	GetAll() ([]models.News, error)
	GetAllByQuery(categoryId, tagId, pageSize, page int) ([]models.News, error)
	GetById(id int) (models.News, error)
	GetAllShortNews() ([]models.ShortNews, error)
	GetCount(categoryId, tagId int) (int, error)
}

type Tags interface {
	GetAll() ([]models.Tags, error)
	GetById(id int) (models.Tags, error)
}

type Categories interface {
	GetAll() ([]models.Categories, error)
	GetById(id int) (models.Categories, error)
}

type Service struct {
	News
	Tags
	Categories
}

func New(db db.DB) *Service {
	return &Service{
		News:       methods.NewNewsService(db.News),
		Tags:       methods.NewTagsService(db.Tags),
		Categories: methods.NewCategoriesService(db.Categories),
	}
}
