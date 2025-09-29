package methods

import (
	"News-portal/internal/db"
	"News-portal/internal/db/models"
)

type NewsService struct {
	db db.News
}

func NewNewsService(db db.News) *NewsService {
	return &NewsService{
		db: db,
	}
}

func (s *NewsService) GetAll() ([]models.News, error) {
	return s.db.GetAll()
}

func (s *NewsService) GetAllByQuery(categoryId, tagId, pageSize, page int) ([]models.News, error) {
	news, err := s.db.GetAllByQuery(categoryId, tagId, pageSize, page)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (s *NewsService) GetById(id int) (models.News, error) {
	return s.db.GetById(id)
}

func (s *NewsService) GetAllShortNews() ([]models.ShortNews, error) {
	return s.db.GetAllShortNews()
}

func (s *NewsService) GetCount(categoryId, tagId int) (int, error) {
	if categoryId != 0 && tagId == 0 {
		return s.db.GetCountByCategory(categoryId)
	} else if categoryId == 0 && tagId != 0 {
		return s.db.GetCountByTag(tagId)
	} else if categoryId != 0 && tagId != 0 {
		return s.db.GetCountByCategoryAndTag(categoryId, tagId)
	} else {
		return s.db.GetCount()
	}
}
