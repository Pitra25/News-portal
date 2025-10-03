package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TagRepo struct {
	db *sqlx.DB
}

func NewTagsPG(db *sqlx.DB) *TagRepo {
	return &TagRepo{
		db: db,
	}
}

func (m *TagRepo) GetAll() ([]Tags, error) {
	var tagsArr []Tags

	query := fmt.Sprint("SELECT * FROM ", tagsTable)

	if err := m.db.Select(&tagsArr, query); err != nil {
		return tagsArr, err
	}

	return tagsArr, nil
}
