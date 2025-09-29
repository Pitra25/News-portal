package db

import (
	"News-portal/internal/db/methods"
	"News-portal/internal/db/models"

	"github.com/jmoiron/sqlx"
)

type News interface {
	GetAll() ([]models.News, error)
	GetById(id int) (models.News, error)
	GetAllShortNews() ([]models.ShortNews, error)
	GetAllByQuery(categoryId, tagId, pageSize, page int) ([]models.News, error)

	GetCountByCategoryAndTag(categoryId, tagId int) (int, error)
	GetCountByCategory(categoryId int) (int, error)
	GetCountByTag(tagId int) (int, error)
	GetCount() (int, error)
}

type Tags interface {
	GetAll() ([]models.Tags, error)
	GetById(id int) (models.Tags, error)
}

type Categories interface {
	GetAll() ([]models.Categories, error)
	GetById(id int) (models.Categories, error)
}

type DB struct {
	News
	Tags
	Categories
}

func New(db *sqlx.DB) *DB {
	return &DB{
		News:       methods.NewNewsPG(db),
		Tags:       methods.NewTagsPG(db),
		Categories: methods.NewCategoriesPG(db),
	}
}
