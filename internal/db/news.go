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

func NewNews(db *sqlx.DB) *NewsRepo {
	return &NewsRepo{
		db: db,
	}
}

func (m *NewsRepo) GetByFilters(fil Filters) ([]News, error) {
	// formation of restrictions
	var limit, offset = fil.Page.paginator()

	query := `SELECT n.*, c."categoryId" FROM ` + newsTable + ` n INNER JOIN ` + categoriesTable + ` c ON c."categoryId" = n."categoryId"
		WHERE ` + fil.News.NewFilters() + ` ORDER BY n."publishedAt" DESC LIMIT :limit OFFSET :offset`

	params := map[string]interface{}{
		"statusID":   newsStatus,
		"categoryID": fil.News.CategoryId,
		"tagID":      fil.News.TagId,
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
		query  = `SELECT * FROM ` + newsTable + ` n WHERE n."newsId" = $1 AND n."statusId" = $2 AND n."publishedAt" <= now()`
		params = []interface{}{id, newsStatus}
	)

	err := m.db.Get(&result, query, params...)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return result, nil
		}
		return result, err
	}

	return result, nil
}

func (m *NewsRepo) GetCount(filter Filters) (int, error) {

	var (
		count  int
		query  = `SELECT COUNT(*) FROM ` + newsTable + ` n WHERE` + filter.News.NewFilters()
		params = map[string]interface{}{
			"statusID":   newsStatus,
			"categoryID": filter.News.CategoryId,
			"tagID":      filter.News.TagId,
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
