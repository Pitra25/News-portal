package db

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type TagsPG struct {
	db *sqlx.DB
}

func NewTagsPG(db *sqlx.DB) *TagsPG {
	return &TagsPG{
		db: db,
	}
}

func (m *TagsPG) GetAll() ([]Tags, error) {
	var tagsArr []Tags

	if err := m.db.Select(
		&tagsArr,
		"SELECT * FROM newsportal.tags",
	); err != nil {
		return tagsArr, err
	}

	return tagsArr, nil
}

func (m *TagsPG) GetById(id int) (Tags, error) {
	var tag Tags
	if err := m.db.Get(
		&tag,
		`SELECT * FROM newsportal.tags WHERE "tagId"=$1`,
		id,
	); err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return tag, nil
		}
		return tag, err
	}

	return tag, nil
}
