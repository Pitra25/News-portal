package method

import (
	"News-portal/internal/db/model"
	"database/sql"
	"fmt"

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

func (m *CategoriesPG) GetAll() ([]model.Categories, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s",
		model.CategoriesTable,
	)

	var categories []model.Categories

	rows, err := m.db.Query(query)
	if err != nil {
		return categories, err
	}

	for rows.Next() {
		var category model.Categories

		if err := rows.Scan(
			&category.CategoryID,
			&category.Title,
			&category.OrderNumber,
			&category.StatusID,
		); err != nil {
			if err == sql.ErrNoRows {
				return categories, nil
			}
			return categories, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (m *CategoriesPG) GetById(id int) (model.Categories, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE categoryId=?",
		model.CategoriesTable,
	)

	var category model.Categories
	if err := m.db.QueryRow(query, id).Scan(
		&category.CategoryID,
		&category.Title,
		&category.OrderNumber,
		&category.StatusID,
	); err != nil {
		if err == sql.ErrNoRows {
			return category, nil
		}
		return category, err
	}

	return category, nil
}
