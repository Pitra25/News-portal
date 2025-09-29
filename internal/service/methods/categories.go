package methods

import (
	"News-portal/internal/db"
	"News-portal/internal/db/models"
)

type CategoriesService struct {
	db db.Categories
}

func NewCategoriesService(db db.Categories) *CategoriesService {
	return &CategoriesService{
		db: db,
	}
}

func (s *CategoriesService) GetAll() ([]models.Categories, error) {
	return s.db.GetAll()
}

func (s *CategoriesService) GetById(id int) (models.Categories, error) {
	return s.db.GetById(id)
}
