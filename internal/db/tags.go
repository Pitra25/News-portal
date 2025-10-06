package db

import (
	"News-portal/output"

	"github.com/go-pg/pg/v10"
)

type TagRepo struct {
	db *pg.DB
}

func NewTags(db *pg.DB) *TagRepo {
	return &TagRepo{
		db: db,
	}
}

func (m *TagRepo) GetAll() ([]output.Tag, error) {
	var tagsArr []output.Tag

	err := m.db.Model(&tagsArr).
		Where(`"statusId" = ?`, newsStatus).
		Select()
	if err != nil {
		return nil, err
	}

	return tagsArr, nil
}

func (m *TagRepo) GetByID(ids []int) ([]output.Tag, error) {
	var tagsArr []output.Tag

	if err := m.db.Model(&tagsArr).
		Where(`"tagId" IN (?)`, pg.In(ids)).
		Where(`"statusId" = ?`, newsStatus).
		Select(); err != nil {
		return nil, err
	}

	return tagsArr, nil
}
