package db

import (
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type NewsPG struct {
	db *sqlx.DB
}

func NewNewsPG(db *sqlx.DB) *NewsPG {
	return &NewsPG{
		db: db,
	}
}

var publishedAt = time.Now()

func (m *NewsPG) GetAll() ([]News, error) {
	var results = []News{}

	if err := m.db.Select(
		&results,
		`SELECT * FROM newsportal.news WHERE "statusId"= :statusId AND "publishedAt" <= :publishedAt`,
		map[string]interface{}{
			"statusId":    newsStatus,
			"publishedAt": publishedAt,
		},
	); err != nil {
		return nil, err
	}

	return results, nil
}

func (m *NewsPG) GetAllByQuery(categoryId, tagId, pageSize, page int) ([]News, error) {
	var limit, offset = 10, 0

	if page == 1 {
		limit = pageSize
		offset = 0
	} else {
		limit = pageSize
		offset = page * pageSize
	}
	// TODO переделать возвращает 1 запись
	//if categoryId > 0 {
	//	query_v2 = query_v2 + fmt.Sprintf(" AND categoryId = %d", categoryId)
	//}
	//if tagId > 0 {
	//	query_v2 = query_v2 + fmt.Sprintf(" AND tagId = %d", tagId)
	//}
	//
	rows, err := m.db.NamedQuery(
		`SELECT
			n.*,
			c."categoryId"
		FROM newsportal.news n
		INNER JOIN newsportal.categories c ON c."categoryId" = n."categoryId"
		WHERE n."statusId" = :statusId
			AND n."publishedAt" <= :publishedAt
			AND n."categoryId" = :categoryId
			AND :tagIds = ANY("tagIds")
		ORDER BY n."publishedAt" DESC
		LIMIT :limit OFFSET :offset`,
		map[string]interface{}{
			"statusId":    newsStatus,
			"publishedAt": publishedAt,
			"categoryId":  categoryId,
			"tagIds":      tagId,
			"limit":       limit,
			"offset":      offset,
		},
	)
	if err != nil {
		return nil, err
	}

	//slog.Info(
	//	"categoryId", categoryId,
	//	"tagId", tagId,
	//	"pageSize", pageSize,
	//	"offset", offset,
	//)

	//query_v2 := fmt.Sprintf("SELECT * FROM newsportal.news WHERE \"statusId\"=%s", newsStatus)
	//if categoryId > 0 {
	//	query_v2 = query_v2 + fmt.Sprintf(" AND categoryId = %d", categoryId)
	//}
	//if tagId > 0 {
	//	query_v2 = query_v2 + fmt.Sprintf(" AND tagId = %d", tagId)
	//}
	//query_v2 = query_v2 + fmt.Sprintf("ORDER BY \"publishedAt\" DESC LIMIT %d OFFSET %d", limit, offset)

	//query :=
	//	`SELECT
	//		*
	//	FROM newsportal.news
	//	WHERE "statusId"=$1
	//		AND "publishedAt" <= $2
	//		AND "categoryId" = $3
	//		AND ($4 = 0 OR $4 = ANY("tagIds"))
	//	ORDER BY "publishedAt" DESC
	//	LIMIT $5 OFFSET $6`
	//rows, err := m.db.Query(
	//	query,
	//	newsStatus,  //$1
	//	publishedAt, //$2
	//	categoryId,  //$3
	//	tagId,       //$4
	//	limit,       //$5
	//	offset,      //$6
	//)

	//slog.Info("query_v2", query_v2)

	//rows, err := m.db.NamedQuery(
	//	`SELECT *
	//	FROM newsportal.news
	//	WHERE "statusId" = :statusId
	//	  	AND "categoryId" = :categoryId
	//		AND "publishedAt" <= :publishedAt
	//		AND  :tagId = ANY("tagIds")
	//	ORDER BY "publishedAt" DESC
	//	LIMIT :pageSize OFFSET :page`,
	//	map[string]interface{}{
	//		"statusId":    newsStatus,
	//		"publishedAt": publishedAt,
	//		"categoryId":  categoryId,
	//		"tagId":       tagId,
	//		"pageSize":    pageSize,
	//		"page":        page,
	//	},
	//)
	//if err != nil {
	//	return nil, err
	//}

	var results []News
	for rows.Next() {
		var news News

		slog.Info("rows:", "rows", rows)
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

func removeDuper(numbers pq.Int64Array) pq.Int64Array {
	seen := make(map[int64]bool)
	result := make([]int64, 0)

	for _, v := range numbers {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func (m *NewsPG) GetById(id int) (News, error) {

	var result News
	if err := m.db.Get(
		&result,
		`SELECT * 
		FROM newsportal.news
		WHERE "newsId" = $1 
			AND "statusId" = $2 
			AND "publishedAt" <= $3`,
		id,
		newsStatus,
		publishedAt,
	); err != nil {
		slog.Error("Error scanning row", "error", err)
		if errors.Is(sql.ErrNoRows, err) {
			return result, nil
		}
		return result, err
	}

	return result, nil
}

func (m *NewsPG) GetAllShortNews() ([]ShortNews, error) {

	var results []ShortNews
	rows, err := m.db.NamedQuery(
		`SELECT 
			"newsId", 
			"title", 
			"categoryId", 
			"tagIds", 
			"publishedAt" 
		FROM newsportal.news
		WHERE "statusId" = :statusId 
			AND "publishedAt" <= :publishedAt
		ORDER BY "publishedAt" DESC`,
		map[string]interface{}{
			"statusId":    newsStatus,
			"publishedAt": publishedAt,
		},
	)
	if err != nil {
		return results, err
	}

	//var results []News
	for rows.Next() {
		var news ShortNews

		if err := rows.Scan(
			&news.NewsID,
			&news.Title,
			&news.CategoryID,
			&news.TagIds,
			&news.PublishedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		news.TagIds = removeDuper(news.TagIds)

		results = append(results, news)
	}

	return results, nil
}

func (m *NewsPG) GetCount(categoryId, tagId int) (int, error) {
	var count int

	if err := m.db.Get(
		&count,
		`SELECT COUNT(*)
		FROM newsportal.news
		WHERE "statusId" = $1 
			AND "publishedAt" <= $2
			AND "categoryId" = $3
			AND ($4 = 0 OR $4 = ANY("tagIds"))`,
		newsStatus,
		publishedAt,
		categoryId,
		tagId,
	); err != nil {
		return 0, err
	}

	return count, nil
}
