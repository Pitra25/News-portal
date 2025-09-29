package methods

import (
	"News-portal/internal/db/models"
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

func (m *CategoriesPG) GetAll() ([]models.Categories, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.CategoriesTable,
	)

	var categories []models.Categories

	rows, err := m.db.Query(query)
	if err != nil {
		return categories, err
	}

	for rows.Next() {
		var category models.Categories

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

func (m *CategoriesPG) GetById(id int) (models.Categories, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE categoryId=?",
		models.CategoriesTable,
	)

	var category models.Categories
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
