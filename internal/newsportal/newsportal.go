package newsportal

import (
	"News-portal/internal/db"
	"log/slog"
)

type Service struct {
	News       *NewsService
	Tags       *TagsService
	Categories *CategoriesService
}

func NewRepo(db *db.DB) *Service {
	slog.Debug("newsportal initialization")
	return &Service{
		News:       NewNewsService(db),
		Tags:       NewTagsService(db),
		Categories: NewCategoriesService(db),
	}
}
