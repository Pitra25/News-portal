package test

import (
	"News-portal/internal/db"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dbo  db.DB
	repo db.NewsRepo
)

func TestMain(m *testing.M) {
	var t *testing.T
	dbo = Setup(t)
	repo = db.NewNewsRepo(dbo)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func intP(i int) *int {
	return &i
}

/*** Category ***/

func Test_AddCategory(t *testing.T) {
	res, clean := Category(t, dbo, nil, WithFakeCategory)
	defer clean()

	fmt.Print("res: ", res, "\n")
}

func Test_CategoriesByFilters(t *testing.T) {
	tags, err := repo.CategoriesByFilters(
		context.Background(),
		nil,
		db.PagerNoLimit,
	)
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(tags), minLength, fmt.Sprint("len: ", len(tags)))
}

func Test_CountCategories(t *testing.T) {
	tests := []struct {
		name    string
		args    *db.CategorySearch
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "not filters",
			args:    nil,
			want:    10,
			wantErr: assert.NoError,
		},
		{
			name: "filter status 2",
			args: &db.CategorySearch{
				StatusID: intP(2),
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "filter status 3",
			args: &db.CategorySearch{
				StatusID: intP(3),
			},
			want:    0,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countNews, err := repo.CountCategories(context.Background(), tt.args)
			assert.NoError(t, err)

			assert.Equal(t, countNews, tt.want, fmt.Sprint("countNews: ", countNews))
		})
	}
}

/*** News ***/

func Test_AddNews(t *testing.T) {
	res, clean := News(t, dbo, nil, WithFakeNews, WithNewsRelations)
	defer clean()

	fmt.Print("res: ", res, "\n")
}

func Test_NewsByID(t *testing.T) {
	tests := []struct {
		name    string
		args    int
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Get news by id (1)",
			args:    1,
			want:    "Иван Петров",
			wantErr: assert.NoError,
		},
		{
			name:    "Get news by id (2)",
			args:    11,
			want:    "Анна Петрова",
			wantErr: assert.NoError,
		},
		{
			name:    "Get news by id (3)",
			args:    12,
			want:    "Михаил Семенов",
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			news, err := repo.NewsByID(context.Background(), tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByFiltersNews() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.NotNil(t, news, fmt.Sprint("no data found"))
			assert.Equal(t, news.Author, tt.want, fmt.Sprint("Author: ", news.Author))
		})
	}
}

func Test_CountNews(t *testing.T) {
	tests := []struct {
		name    string
		args    *db.NewsSearch
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "not filters",
			args:    nil,
			want:    26,
			wantErr: assert.NoError,
		},
		{
			name: "filters status",
			args: &db.NewsSearch{
				StatusID: intP(2),
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "Get news by id (3)",
			args: &db.NewsSearch{
				CategoryID: intP(3),
			},
			want:    2,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countNews, err := repo.CountNews(context.Background(), tt.args)
			assert.NoError(t, err)

			assert.Equal(t, countNews, tt.want, fmt.Sprint("countNews: ", countNews))
		})
	}
}

/*** Tag ***/

func Test_AddTag(t *testing.T) {
	res, clean := Tag(t, dbo, nil, WithFakeTag)
	defer clean()

	fmt.Print("res: ", res, "\n")
}

func Test_CountTags(t *testing.T) {
	tests := []struct {
		name    string
		args    *db.TagSearch
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "not filters",
			args:    nil,
			want:    10,
			wantErr: assert.NoError,
		},
		{
			name: "filters status 5",
			args: &db.TagSearch{
				StatusID: intP(5),
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "filters status 8",
			args: &db.TagSearch{
				StatusID: intP(8),
			},
			want:    0,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countNews, err := repo.CountTags(context.Background(), tt.args)
			assert.NoError(t, err)

			assert.Equal(t, countNews, tt.want, fmt.Sprint("countNews: ", countNews))
		})
	}
}
