package db

import (
	"github.com/go-pg/pg/v10"
)

type CategoryRepo struct {
	db *pg.DB
}

func NewCategory(db *pg.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (m *CategoryRepo) GetAll() ([]Category, error) {
	var arrCategories []Category

	err := filStatus(m.db.Model(&arrCategories)).
		Select()
	if err != nil {
		return nil, err
	}

	return arrCategories, nil
}

func (m *CategoryRepo) GetById(ids []int) ([]Category, error) {
	var result []Category

	if err := filStatus(m.db.Model(&result)).
		Where(`"t"."categoryId" IN (?)`, pg.In(ids)).
		Select(); err != nil {
		return nil, err
	}

	return result, nil
}
