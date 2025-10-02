package newsportal

import (
	"News-portal/internal/db"
)

func categoriesDtoToJson(c db.Categories) Categories {
	return Categories{
		CategoryID:  c.CategoryID,
		Title:       c.Title,
		OrderNumber: c.OrderNumber,
	}
}
