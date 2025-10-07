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

func (m *TagRepo) GetAll() ([]Tag, error) {
	var tagsArr []Tag

	err := filStatus(m.db.Model(&tagsArr)).
		Select()
	if err != nil {
		return nil, err
	}

	return tagsArr, nil
}

func (m *TagRepo) GetByID(ids []int) ([]Tag, error) {
	var tagsArr []Tag

	if err := filStatus(m.db.Model(&tagsArr)).
		Where(`"t"."tagId" IN (?)`, pg.In(ids)).
		Select(); err != nil {
		return nil, err
	}

	return tagsArr, nil
}
