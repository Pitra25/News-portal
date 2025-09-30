package method

import (
	"News-portal/internal/db"
	"News-portal/internal/service/model"
)

type CategoriesService struct {
	db *db.DB
}

func NewCategoriesService(db *db.DB) *CategoriesService {
	return &CategoriesService{
		db: db,
	}
}

func (s *CategoriesService) GetAll() ([]model.Categories, error) {
	categories := []model.Categories{}

	categoriesArr, err := s.db.Categories.GetAll()
	if err != nil {
		return categories, err
	}

	for _, v := range categoriesArr {
		category := model.Categories{
			CategoryID:  v.CategoryID,
			Title:       v.Title,
			OrderNumber: v.OrderNumber,
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (s *CategoriesService) GetById(id int) (model.Categories, error) {
	category, err := s.db.Categories.GetById(id)
	if err != nil {
		return model.Categories{}, err
	}

	return model.Categories{
		CategoryID:  category.CategoryID,
		Title:       category.Title,
		OrderNumber: category.OrderNumber,
	}, nil
}
