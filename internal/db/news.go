package db

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type NewsRepo struct {
	db *sqlx.DB
}

func NewNewsPG(db *sqlx.DB) *NewsRepo {
	return &NewsRepo{
		db: db,
	}
}

func (m *NewsRepo) GetByFilters(filter FiltersNews, pageF PageFilters) ([]News, error) {
	// default values
	var limit, offset = 10, 0

	if pageF.Page == 1 {
		limit = pageF.PageSize
		offset = 0
	} else {
		limit = pageF.PageSize
		offset = pageF.Page * pageF.PageSize
	}

	// forming a request and parameters
	//query := `SELECT n.*, n."statusId", c."categoryId" FROM newsportal.news n INNER JOIN newsportal.categories c ON c."categoryId" = n."categoryId"
	//		WHERE` + filter.NewFilters() + ` LIMIT :limit OFFSET :offset`

	query := `SELECT n.*, c."categoryId" FROM newsportal.news n INNER JOIN newsportal.categories c ON c."categoryId" = n."categoryId"
		WHERE ` + filter.NewNewsFilters() + ` ORDER BY n."publishedAt" DESC LIMIT :limit OFFSET :offset`

	params := map[string]interface{}{
		"statusID":   newsStatus,
		"categoryID": filter.CategoryId,
		"tagID":      filter.TagId,
		"limit":      limit,
		"offset":     offset,
	}

	// query execution
	rows, err := m.db.NamedQuery(query, params)
	if err != nil {
		return nil, err
	}

	// forming a news list
	var results []News
	for rows.Next() {
		var news News

		if err := rows.StructScan(&news); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		news.TagIds = removeDuper(news.TagIds)

		results = append(results, news)
	}
	slog.Info("w", "results", results)

	return results, nil
}

func (m *NewsRepo) GetById(id int) (News, error) {

	var (
		result News
		query  = `SELECT * FROM newsportal.news n WHERE n."newsId" = $1 AND n."statusId" = $2 AND n."publishedAt" <= now()`
		params = []interface{}{id, newsStatus}
	)

	err := m.db.Get(&result, query, params...)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return result, nil
		}
		return result, err
	}

	//for rows.Next() {
	//	if err := rows.Scan(&result); err != nil {
	//		return result, err
	//	}
	//}

	return result, nil
}

func (m *NewsRepo) GetCount(filter FiltersNews) (int, error) {

	var (
		count  int
		query  = `SELECT COUNT(*) FROM newsportal.news WHERE` + filter.NewFilters()
		params = map[string]interface{}{
			"statusID":   newsStatus,
			"categoryID": filter.CategoryId,
			"tagID":      filter.TagId,
		}
	)

	rows, err := m.db.NamedQuery(query, params)
	if err != nil {
		return count, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return count, err
		}
	}

	return count, nil
}
