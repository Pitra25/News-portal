package methods

import (
	"News-portal/internal/db/models"
	"database/sql"
	"fmt"

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

func (m *TagsPG) GetAll() ([]models.Tags, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s",
		models.TagsTable,
	)

	var tags []models.Tags

	rows, err := m.db.Query(query)
	if err != nil {
		return tags, err
	}

	for rows.Next() {
		var tag models.Tags

		if err := rows.Scan(
			&tag.TagID,
			&tag.Title,
			&tag.StatusID,
		); err != nil {
			if err == sql.ErrNoRows {
				return tags, nil
			}
			return tags, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (m *TagsPG) GetById(id int) (models.Tags, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE tagId=?",
		models.TagsTable,
	)

	var tag models.Tags
	if err := m.db.QueryRow(query, id).Scan(
		&tag.TagID,
		&tag.Title,
		&tag.StatusID,
	); err != nil {
		if err == sql.ErrNoRows {
			return tag, nil
		}
		return tag, err
	}

	return tag, nil
}
