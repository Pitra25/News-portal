package methods

import (
	"News-portal/internal/db"
	"News-portal/internal/db/models"
)

type TagsService struct {
	db db.Tags
}

func NewTagsService(db db.Tags) *TagsService {
	return &TagsService{
		db: db,
	}
}

func (s *TagsService) GetAll() ([]models.Tags, error) {
	return s.db.GetAll()
}

func (s *TagsService) GetById(id int) (models.Tags, error) {
	return s.db.GetById(id)
}
