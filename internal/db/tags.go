package db

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type TagRepo struct {
	db *sqlx.DB
}

func TagsInit(db *sqlx.DB) *TagRepo {
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

func (m *TagRepo) GetByID(ids []int) ([]Tags, error) {
	var tagsArr []Tags

	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))

	for i, id := range ids {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	query := fmt.Sprintf(`SELECT * FROM %s t WHERE t."tagId" IN (%s)`,
		tagsTable, strings.Join(placeholders, ","))

	if err := m.db.Select(&tagsArr, query, args...); err != nil {
		return tagsArr, err
	}

	return tagsArr, nil
}
