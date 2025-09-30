package method

import (
	"News-portal/internal/db"
	"News-portal/internal/service/model"
	"log/slog"
)

type NewsService struct {
	db *db.DB
}

func NewNewsService(db *db.DB) *NewsService {
	return &NewsService{
		db: db,
	}
}

func (s *NewsService) GetAll() ([]model.News, error) {
	respNews := []model.News{}
	news, err := s.db.News.GetAll()
	if err != nil {
		return respNews, err
	}

	for _, v := range news {
		category, err := s.db.Categories.GetById(v.CategoryID)
		if err != nil {
			slog.Error("get title category by id: ", v.CategoryID)
		}

		tags := []model.Tag{}
		for _, v := range v.TagIds {
			tag, err := s.db.Tags.GetById(v)
			if err != nil {
				slog.Error("get tag by id: ", v)
				continue
			}
			tags = append(tags, model.Tag{
				TagID: tag.TagID,
				Title: tag.Title,
			})
		}

		newsV := model.News{
			NewsID:  v.NewsID,
			Title:   v.Title,
			Content: v.Content,
			Author:  v.Author,
			CategoryID: model.Category{
				CategoryID: v.CategoryID,
				Title:      category.Title,
			},
			TagIds:      tags,
			PublishedAt: v.PublishedAt,
		}

		respNews = append(respNews, newsV)
	}

	return respNews, nil
}

func (s *NewsService) GetAllByQuery(categoryId, tagId, pageSize, page int) ([]model.News, error) {
	respNews := []model.News{}

	news, err := s.db.News.GetAllByQuery(categoryId, tagId, pageSize, page)
	if err != nil {
		return respNews, err
	}

	for _, v := range news {
		category, err := s.db.Categories.GetById(v.CategoryID)
		if err != nil {
			slog.Error("get title category by id: ", v.CategoryID)
		}

		tags := []model.Tag{}
		for _, v := range v.TagIds {
			tag, err := s.db.Tags.GetById(v)
			if err != nil {
				slog.Error("get tag by id: ", v)
				continue
			}
			tags = append(tags, model.Tag{
				TagID: tag.TagID,
				Title: tag.Title,
			})
		}

		newsV := model.News{
			NewsID:  v.NewsID,
			Title:   v.Title,
			Content: v.Content,
			Author:  v.Author,
			CategoryID: model.Category{
				CategoryID: v.CategoryID,
				Title:      category.Title,
			},
			TagIds:      tags,
			PublishedAt: v.PublishedAt,
		}

		respNews = append(respNews, newsV)
	}

	return respNews, nil
}

func (s *NewsService) GetById(id int) (model.News, error) {

	news, err := s.db.News.GetById(id)
	if err != nil {
		return model.News{}, err
	}

	category, err := s.db.Categories.GetById(news.CategoryID)
	if err != nil {
		slog.Error("get title category by id: ", news.CategoryID)
	}

	tags := []model.Tag{}
	for _, v := range news.TagIds {
		tag, err := s.db.Tags.GetById(v)
		if err != nil {
			slog.Error("get tag by id: ", v)
			continue
		}
		tags = append(tags, model.Tag{
			TagID: tag.TagID,
			Title: tag.Title,
		})
	}

	return model.News{
		NewsID:  news.NewsID,
		Title:   news.Title,
		Content: news.Content,
		Author:  news.Author,
		CategoryID: model.Category{
			CategoryID: news.CategoryID,
			Title:      category.Title,
		},
		TagIds:      tags,
		PublishedAt: news.PublishedAt,
	}, nil

}

func (s *NewsService) GetAllShortNews() ([]model.ShortNews, error) {

	respNews := []model.ShortNews{}

	news, err := s.db.News.GetAllShortNews()
	if err != nil {
		return respNews, err
	}

	for _, v := range news {
		category, err := s.db.Categories.GetById(v.CategoryID)
		if err != nil {
			slog.Error("get title category by id: ", v.CategoryID)
		}

		tags := []model.Tag{}
		for _, v := range v.TagIds {
			tag, err := s.db.Tags.GetById(v)
			if err != nil {
				slog.Error("get tag by id: ", v)
				continue
			}
			tags = append(tags, model.Tag{
				TagID: tag.TagID,
				Title: tag.Title,
			})
		}

		newsV := model.ShortNews{
			NewsID: v.NewsID,
			Title:  v.Title,
			CategoryID: model.Category{
				CategoryID: v.CategoryID,
				Title:      category.Title,
			},
			TagIds:      tags,
			PublishedAt: v.PublishedAt,
		}

		respNews = append(respNews, newsV)
	}

	return respNews, nil
}

func (s *NewsService) GetCount(categoryId, tagId int) (int, error) {
	if categoryId != 0 && tagId == 0 {
		return s.db.News.GetCountByCategory(categoryId)
	} else if categoryId == 0 && tagId != 0 {
		return s.db.News.GetCountByTag(tagId)
	} else if categoryId != 0 && tagId != 0 {
		return s.db.News.GetCountByCategoryAndTag(categoryId, tagId)
	} else {
		return s.db.News.GetCount()
	}
}
