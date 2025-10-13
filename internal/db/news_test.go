package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Repo_GetNewsByFilters(t *testing.T) {
	tests := []struct {
		name    string
		args    Filters
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get all news by query (1)",
			args: Filters{
				CategoryId: 2,
				TagId:      1,
				PageSize:   10,
				Page:       1,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (2)",
			args: Filters{
				CategoryId: 4,
				TagId:      1,
				PageSize:   10,
				Page:       1,
			},
			want:    4,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (3)",
			args: Filters{
				CategoryId: 5,
				TagId:      1,
				PageSize:   10,
				Page:       1,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (4)",
			args: Filters{
				CategoryId: 10,
				TagId:      1,
				PageSize:   10,
				Page:       1,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (4)",
			args: Filters{
				CategoryId: 10,
				TagId:      0,
				PageSize:   10,
				Page:       1,
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (4)",
			args: Filters{
				CategoryId: 0,
				TagId:      1,
				PageSize:   10,
				Page:       1,
			},
			want:    10,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Repo{
				db: connDB,
			}
			list, err := m.GetNewsByFilters(context.Background(), tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByFiltersNews() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.Len(t, list, tt.want, fmt.Sprint("len: ", len(list)))
		})
	}
}

func Test_Repo_GetNewsById(t *testing.T) {
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

			m := &Repo{
				db: connDB,
			}
			news, err := m.GetNewsById(context.Background(), tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByFiltersNews() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.NotNil(t, news, fmt.Sprint("no data found"))
			assert.Equal(t, news.Author, tt.want, fmt.Sprint("Author: ", news.Author))
		})
	}
}

func Test_Repo_GetNewsCount(t *testing.T) {
	tests := []struct {
		name    string
		args    Filters
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get count by categoryId 1 and tagId 0",
			args: Filters{
				CategoryId: 1,
				TagId:      0,
			},
			want:    3,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 0",
			args: Filters{
				CategoryId: 2,
				TagId:      0,
			},
			want:    3,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 1 and tagId 1",
			args: Filters{
				CategoryId: 1,
				TagId:      1,
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 1",
			args: Filters{
				CategoryId: 2,
				TagId:      1,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 2",
			args: Filters{
				CategoryId: 2,
				TagId:      2,
			},
			want:    3,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 7 and tagId 1",
			args: Filters{
				CategoryId: 7,
				TagId:      1,
			},
			want:    4,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &Repo{
				db: connDB,
			}

			count, err := m.GetNewsCount(context.Background(), tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("Count() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.Equal(t, tt.want, count, fmt.Sprint("Count() count: ", count))
		})
	}
}

func Test_Repo_GetListCategories(t *testing.T) {
	m := &Repo{
		db: connDB,
	}
	tags, err := m.GetListCategories(context.Background())
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(tags), minLength, fmt.Sprint("len: ", len(tags)))
}

func Test_Repo_GetTagByID(t *testing.T) {

	tests := []struct {
		name    string
		args    []int
		want    []Tag
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test 1",
			args: []int{1, 2, 10},
			want: []Tag{
				{ID: 1, Title: "Новости", StatusID: 1},
				{ID: 2, Title: "Аналитика", StatusID: 1},
				{ID: 10, Title: "Прогноз", StatusID: 1},
			},
			wantErr: assert.NoError,
		},
		{
			name: "test 2",
			args: []int{1, 6, 9},
			want: []Tag{
				{ID: 1, Title: "Новости", StatusID: 1},
				{ID: 6, Title: "Обзор", StatusID: 1},
				{ID: 9, Title: "Статистика", StatusID: 1},
			},
			wantErr: assert.NoError,
		},
		{
			name: "test 3",
			args: []int{1, 5, 9},
			want: []Tag{
				{ID: 1, Title: "Новости", StatusID: 1},
				{ID: 5, Title: "Реportаж", StatusID: 1},
				{ID: 9, Title: "Статистика", StatusID: 1},
			},
			wantErr: assert.NoError,
		},
		{
			name: "test 4",
			args: []int{1, 4, 5},
			want: []Tag{
				{ID: 1, Title: "Новости", StatusID: 1},
				{ID: 4, Title: "Интервью", StatusID: 1},
				{ID: 5, Title: "Реportаж", StatusID: 1},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Repo{
				db: connDB,
			}
			got, err := m.GetTagByID(context.Background(), tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetTagByID(%v)", tt.args)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetTagByID(%v)", tt.args)
		})
	}
}
