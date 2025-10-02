package newsportal

import (
	"News-portal/internal/db"
)

type CategoriesService struct {
	db *db.DB
}

func NewCategoriesService(db *db.DB) *CategoriesService {
	return &CategoriesService{
		db: db,
	}
}

func (s *CategoriesService) GetAll() ([]Categories, error) {
	categories := []Categories{}

	categoriesArr, err := s.db.Categories.GetAll()
	if err != nil {
		return categories, err
	}

	for _, v := range categoriesArr {
		categories = append(
			categories,
			categoriesDtoToJson(v),
		)
	}

	return categories, nil
}

func (s *CategoriesService) GetById(id int) (Categories, error) {
	category, err := s.db.Categories.GetById(id)
	if err != nil {
		return Categories{}, err
	}

	return categoriesDtoToJson(category), nil
}
