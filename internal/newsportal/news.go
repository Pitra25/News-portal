package newsportal

import (
	"News-portal/internal/db"
	"log/slog"

	"github.com/lib/pq"
)

type NewsService struct {
	db *db.DB
}

func NewNewsService(db *db.DB) *NewsService {
	return &NewsService{
		db: db,
	}
}

func (s *NewsService) GetAll() ([]News, error) {
	respNews := []News{}
	newsArr, err := s.db.News.GetAll()
	if err != nil {
		return respNews, err
	}

	for _, v := range newsArr {
		category, err := s.db.Categories.GetById(v.CategoryID)
		if err != nil {
			slog.Error("fail get title category", "id", v.CategoryID)
		}

		respNews = append(
			respNews,
			newsDtoToJson(
				v,
				getTags(s.db, v.TagIds),
				category,
			),
		)
	}

	return respNews, nil
}

func (s *NewsService) GetAllByQuery(categoryId, tagId, pageSize, page int) ([]News, error) {
	respNews := []News{}

	news, err := s.db.News.GetByFilters(categoryId, tagId, pageSize, page)
	if err != nil {
		return respNews, err
	}

	for _, v := range news {
		category, err := s.db.Categories.GetById(v.CategoryID)
		if err != nil {
			slog.Error("fail get title category", "id", v.CategoryID)
		}

		respNews = append(
			respNews,
			newsDtoToJson(
				v,
				getTags(s.db, v.TagIds),
				category,
			),
		)
	}

	return respNews, nil
}

func (s *NewsService) GetById(id int) (News, error) {
	respNews, err := s.db.News.GetById(id)
	if err != nil {
		return News{}, err
	}

	category, err := s.db.Categories.GetById(respNews.CategoryID)
	if err != nil {
		slog.Error("fail get title category", "id", respNews.CategoryID)
	}

	return newsDtoToJson(
		respNews,
		getTags(s.db, respNews.TagIds),
		category,
	), nil
}

func (s *NewsService) GetAllShortNews() ([]ShortNews, error) {

	respNews := []ShortNews{}

	news, err := s.db.News.GetAllShortNews()
	if err != nil {
		return respNews, err
	}

	for _, v := range news {
		category, err := s.db.Categories.GetById(v.CategoryID)
		if err != nil {
			slog.Error("fail get title category", "id", v.CategoryID)
		}

		respNews = append(
			respNews,
			shortNewsDtoToJson(
				v,
				getTags(s.db, v.TagIds),
				category,
			),
		)
	}

	return respNews, nil
}

func (s *NewsService) GetCount(categoryId, tagId int) (int, error) {
	return s.db.News.GetCount(categoryId, tagId)
}

func getTags(db *db.DB, tagsId pq.Int64Array) []Tag {
	result := []Tag{}

	for _, v := range tagsId {
		tag, err := db.Tags.GetById(int(v))
		if err != nil {
			slog.Error("fail get tag", "id", v)
			continue
		}
		result = append(
			result,
			tagDtoToJson(tag),
		)
	}

	return result
}
