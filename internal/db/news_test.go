package db

import (
	"News-portal/internal/db/test"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func intP(i int) *int {
	return &i
}

/*** Category ***/

func Test_AddCategory(t *testing.T) {
	dbo := test.Setup(t)

	res, clean := test.Category(t, dbo, nil, test.WithFakeCategory)
	defer clean()

	fmt.Print(res)

	assert.Len(t, res, 1)
}

func Test_CategoriesByFilters_List(t *testing.T) {
	dbo := test.Setup(t)
	repo := NewNewsRepo(dbo)

	tags, err := repo.CategoriesByFilters(
		context.Background(),
		nil,
		PagerNoLimit,
	)
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(tags), minLength, fmt.Sprint("len: ", len(tags)))
}

/*** News ***/

func Test_AddNews(t *testing.T) {
	dbo := test.Setup(t)

	res, clean := test.News(t, dbo, nil, test.WithFakeNews, test.WithNewsRelations)
	defer clean()

	assert.Len(t, res, 1)
}

func Test_Repo_GetNewsById(t *testing.T) {
	dbo := test.Setup(t)
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
			repo := NewNewsRepo(dbo)

			news, err := repo.NewsByID(context.Background(), tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByFiltersNews() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.NotNil(t, news, fmt.Sprint("no data found"))
			assert.Equal(t, news.Author, tt.want, fmt.Sprint("Author: ", news.Author))
		})
	}
}

func Test_Repo_GetListCategories(t *testing.T) {
	dbo := test.Setup(t)

	tests := []struct {
		name    string
		args    *NewsSearch
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "not filters",
			args:    nil,
			want:    24,
			wantErr: assert.NoError,
		},
		{
			name: "filters status",
			args: &NewsSearch{
				StatusID: intP(2),
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "Get news by id (3)",
			args: &NewsSearch{
				CategoryID: intP(3),
			},
			want:    5,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewNewsRepo(dbo)

			countNews, err := repo.CountNews(context.Background(), tt.args)
			assert.NoError(t, err)

			assert.Equal(t, countNews, tt.want, fmt.Sprint("countNews: ", countNews))
		})
	}
}

/*** Tag ***/

func Test_AddTag(t *testing.T) {
	dbo := test.Setup(t)

	res, clean := test.Tag(t, dbo, nil, test.WithFakeTag)
	defer clean()

	fmt.Print(res)
}

func Test_CountTags(t *testing.T) {
	dbo := test.Setup(t)

	res, clean := test.Tag(t, dbo, nil)
	defer clean()

	assert.Len(t, res, 1)
}
