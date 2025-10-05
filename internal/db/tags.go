package db

import (
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

func (m *TagRepo) GetAll() ([]Tags, error) {
	var tagsArr []Tags

	err := m.db.Model(&tagsArr).
		Where(`"statusId" = ?`, newsStatus).
		Select()
	if err != nil {
		return nil, err
	}

	return tagsArr, nil
}

func (m *TagRepo) GetByID(ids []int64) ([]Tags, error) {
	var tagsArr []Tags

	if err := m.db.Model(&tagsArr).
		Where(`"tagId" IN (?)`, pg.In(ids)).
		Where(`"statusId" = ?`, newsStatus).
		Select(); err != nil {
		return nil, err
	}

	return tagsArr, nil
}
