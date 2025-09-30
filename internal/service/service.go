package service

import (
	"News-portal/internal/db"
	"News-portal/internal/service/method"
	"log/slog"
)

type Service struct {
	News       *method.NewsService
	Tags       *method.TagsService
	Categories *method.CategoriesService
}

func New(db *db.DB) *Service {
	slog.Debug("service initialization")
	return &Service{
		News:       method.NewNewsService(db),
		Tags:       method.NewTagsService(db),
		Categories: method.NewCategoriesService(db),
	}
}
