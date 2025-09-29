package methods

import (
	"News-portal/internal/db/models"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
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

func (m *NewsPG) GetAll() ([]models.News, error) {
	// query := fmt.Sprintf(
	// 	"SELECT n.newsId, n.title, n.content, n.author, c.title, n.tagIds, n.publishedAt "+
	// 		"FROM %s n"+
	// 		"INNER JOIN %s c ON c.categoryId = n.categoryId"+
	// 		"WHERE n.statusId=%b AND n.publishedAt<=%s",
	// 	models.NewsTable, models.CategoriesTable,
	// 	models.NewsStatus, publishedAt,
	// )

	query := fmt.Sprintf(
		"SELECT * FROM %s n "+
			"WHERE statusId=%b AND publishedAt<=%s",
		models.NewsTable,
		models.NewsStatus, publishedAt,
	)

	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}

	var results []models.News
	for rows.Next() {
		var news models.News

		if err := rows.Scan(
			&news.NewsID,
			&news.Title,
			&news.Author,
			&news.CategoryID,
			&news.TagIds,
			&news.CreatedAt,
			&news.PublishedAt,
			&news.StatusID,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		results = append(results, news)
	}

	return results, nil
}

func (m *NewsPG) GetAllByQuery(categoryId, tagId, pageSize, page int) ([]models.News, error) {
	var limit, offset = 0, 0

	if page == 1 {
		offset = 0
	} else {
		limit = pageSize
		offset = page * pageSize
	}

	query :=
		`SELECT 
			n.*, 
			c.categoryId, 
			s.statusId 
		FROM ` + models.NewsTable + `n 
		INNER JOIN` + models.CategoriesTable + `c ON c.categoryId = n.categoryId
		INNER JOIN` + models.StatusesTable + ` s ON s.statusId = n.statusId
		WHERE n.statusId=$1 
			AND n.publishedAt<=$2
			AND n.categoryId = $3 
			AND $4 = ANY(n.tagIds)
		ORDER BY n.publishedAt DESC
		LIMIT $5 OFFSET $6`

	rows, err := m.db.Query(
		query,
		models.NewsStatus, //$1
		publishedAt,       //$2
		categoryId,        //$3
		tagId,             //$4
		limit,             //$5
		offset,            //$6
	)
	if err != nil {
		return nil, err
	}

	var results []models.News
	for rows.Next() {
		var news models.News

		if err := rows.Scan(
			&news.NewsID,
			&news.Title,
			&news.Author,
			&news.CategoryID,
			&news.TagIds,
			&news.CreatedAt,
			&news.PublishedAt,
			&news.StatusID,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		news.TagIds = removeDuper(news.TagIds)

		results = append(results, news)
	}

	// result := clearItem(results)

	return results, nil
}

// func clearItem(news []models.News) []models.News {
// 	for i := range news {
// 		news[i].TagIds = removeDuper(news[i].TagIds)
// 	}
// 	return news
// }

func removeDuper(numbers []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0)

	for _, v := range numbers {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func (m *NewsPG) GetById(id int) (models.News, error) {
	query := fmt.Sprintf(
		"SELECT * FROM %s WHERE newsId=$1 AND statusId=%b AND publishedAt<=%s",
		models.NewsTable, models.NewsStatus, publishedAt,
	)

	var news models.News
	if err := m.db.QueryRow(query, id).Scan(
		&news.NewsID,
		&news.Title,
		&news.Author,
		&news.CategoryID,
		&news.TagIds,
		&news.CreatedAt,
		&news.PublishedAt,
		&news.StatusID,
	); err != nil {
		if err == sql.ErrNoRows {
			return news, nil
		}
		return news, err
	}

	return news, nil
}

func (m *NewsPG) GetAllShortNews() ([]models.ShortNews, error) {
	query := fmt.Sprintf(
		"SELECT newId, title, categoryId, tagIds, publishedAt "+
			"FROM %s WHERE statusId=%b AND publishedAt<=%s",
		models.NewsTable, models.NewsStatus, publishedAt,
	)

	var results []models.ShortNews

	rows, err := m.db.Query(query)
	if err != nil {
		return results, err
	}

	for rows.Next() {
		var news models.ShortNews

		if err := rows.Scan(
			&news.NewsID,
			&news.Title,
			&news.CategoryID,
			&news.TagIds,
			&news.PublishedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				return results, nil
			}
			return results, err
		}
		results = append(results, news)
	}

	return results, nil
}

func (m *NewsPG) GetCountByCategoryAndTag(categoryId, tagId int) (int, error) {

	var count int
	if err := m.db.QueryRow(
		"SELECT COUNT(*) FROM "+models.NewsTable+" "+
			"WHERE categoryId = $1 AND $2 = ANY(tagIds) AND statusId=$3 AND publishedAt<=%4",
		categoryId, tagId, models.NewsStatus, publishedAt).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (m *NewsPG) GetCountByCategory(categoryId int) (int, error) {

	var count int
	if err := m.db.QueryRow(
		"SELECT COUNT(*) FROM "+models.NewsTable+" "+
			"WHERE categoryId = $1 AND statusId=$3 AND publishedAt<=%4",
		categoryId, models.NewsStatus, publishedAt).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (m *NewsPG) GetCountByTag(tagId int) (int, error) {

	var count int
	if err := m.db.QueryRow(
		"SELECT COUNT(*) FROM "+models.NewsTable+" "+
			"WHERE $2 = ANY(tagIds) AND statusId=$3 AND publishedAt<=%4",
		tagId, models.NewsStatus, publishedAt).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (m *NewsPG) GetCount() (int, error) {

	var count int
	if err := m.db.QueryRow(
		"SELECT COUNT(*) FROM "+models.NewsTable+" "+
			"WHERE AND statusId=$3 AND publishedAt<=%4",
		models.NewsStatus, publishedAt).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
