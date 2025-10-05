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

func (m *CategoryRepo) GetAll() ([]Categories, error) {
	var arrCategories []Categories

	err := m.db.Model(&arrCategories).
		Where(`"statusId" = ?`, newsStatus).
		Select()
	if err != nil {
		return nil, err
	}

	return arrCategories, nil
}

func (m *CategoryRepo) GetById(ids []int) ([]Categories, error) {
	var result []Categories

	if err := m.db.Model(&result).
		Where(`"categoryId" IN (?)`, pg.In(ids)).
		Where(`"statusId" = ?`, newsStatus).
		Select(); err != nil {
		return nil, err
	}

	return result, nil
}
