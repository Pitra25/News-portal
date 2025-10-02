package newsportal

import (
	"News-portal/internal/db"
)

func tagDtoToJson(tagDB db.Tags) Tag {
	return Tag{
		TagID: tagDB.TagID,
		Title: tagDB.Title,
	}
}
