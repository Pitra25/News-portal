package db

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type CategoryRepo struct {
	db *sqlx.DB
}

func CategoryInit(db *sqlx.DB) *CategoryRepo {
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

func (m *CategoryRepo) GetById(ids []int) ([]Categories, error) {
	var result []Categories

	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))

	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`SELECT * FROM %s c WHERE c."categoryId" IN (%s)`,
		categoriesTable, strings.Join(placeholders, ","))

	//query := fmt.Sprint(`SELECT * FROM `, categoriesTable, ` WHERE "categoryId" = `, id)
	if err := m.db.Select(&result, query, args...); err != nil {
		return result, err
	}

	return result, nil
}
