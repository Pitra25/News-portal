package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type CategoryRepo struct {
	db *sqlx.DB
}

func NewCategoriesPG(db *sqlx.DB) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (m *CategoryRepo) GetAll() ([]Categories, error) {
	var arrCategories []Categories

	query := fmt.Sprint("SELECT * FROM ", categoriesTable)

	if err := m.db.Select(&arrCategories, query); err != nil {
		return arrCategories, err
	}

	return arrCategories, nil
}
