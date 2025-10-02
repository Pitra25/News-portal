package newsportal

import (
	"News-portal/internal/db"
)

type TagsService struct {
	db *db.DB
}

func NewTagsService(db *db.DB) *TagsService {
	return &TagsService{
		db: db,
	}
}

func (s *TagsService) GetAll() ([]Tag, error) {
	tagsResp := []Tag{}

	tagsArr, err := s.db.Tags.GetAll()
	if err != nil {
		return tagsResp, err
	}

	for _, v := range tagsArr {
		tagsResp = append(
			tagsResp,
			tagDtoToJson(v),
		)
	}

	return tagsResp, nil
}

func (s *TagsService) GetById(id int) (Tag, error) {
	tagResp, err := s.db.Tags.GetById(id)
	if err != nil {
		return Tag{}, err
	}

	return tagDtoToJson(tagResp), nil
}
