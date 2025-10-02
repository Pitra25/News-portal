package db

import (
	"github.com/jmoiron/sqlx"
)

type CategoriesPG struct {
	db *sqlx.DB
}

func NewCategoriesPG(db *sqlx.DB) *CategoriesPG {
	return &CategoriesPG{
		db: db,
	}
}

func (m *CategoriesPG) GetAll() ([]Categories, error) {
	var arrCategories []Categories

	if err := m.db.Select(
		&arrCategories,
		`SELECT * FROM newsportal.categories`,
	); err != nil {
		return arrCategories, err
	}

	return arrCategories, nil
}

func (m *CategoriesPG) GetById(id int) (Categories, error) {
	var category Categories
	if err := m.db.Get(
		&category,
		`SELECT * FROM newsportal.categories WHERE "categoryId" = $1`,
		id,
	); err != nil {
		return category, err
	}

	return category, nil
}
