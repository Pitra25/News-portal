package method

import (
	"News-portal/internal/db"
	"News-portal/internal/service/model"
)

type TagsService struct {
	db *db.DB
}

func NewTagsService(db *db.DB) *TagsService {
	return &TagsService{
		db: db,
	}
}

func (s *TagsService) GetAll() ([]model.Tag, error) {
	tags := []model.Tag{}

	tagsArr, err := s.db.Tags.GetAll()
	if err != nil {
		return tags, err
	}

	for _, v := range tagsArr {
		tag := model.Tag{
			TagID: v.TagID,
			Title: v.Title,
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

func (s *TagsService) GetById(id int) (model.Tag, error) {
	tag, err := s.db.Tags.GetById(id)
	if err != nil {
		return model.Tag{}, err
	}

	return model.Tag{
		TagID: tag.TagID,
		Title: tag.Title,
	}, nil
}
