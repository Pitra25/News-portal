package newsportal

import (
	"News-portal/internal/db"
	"News-portal/internal/db/test"
	"context"
	"fmt"
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var connDB db.NewsRepo

func TestMain(m *testing.M) {
	var t *testing.T
	dbo := test.Setup(t)
	connDB = db.NewNewsRepo(dbo)

	exitCode := m.Run()
	os.Exit(exitCode)
}
func strP(s string) *string {
	return &s
}
func intP(i int) *int {
	return &i
}

func TestManager_GetAllCategory(t *testing.T) {
	m := &Manager{
		repo: connDB,
	}
	categories, err := m.GetAllCategory(context.Background())
	assert.NoError(t, err)

	const minLength = 5
	slog.Info("categories", "cat: ", categories)
	assert.GreaterOrEqual(
		t, len(categories), minLength,
		fmt.Sprint("len: ", len(categories)),
	)
}

func TestManager_GetAllTag(t *testing.T) {
	m := &Manager{
		repo: connDB,
	}
	tags, err := m.GetAllTag(context.Background())
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(tags), minLength, fmt.Sprint("len: ", len(tags)))

}

func TestManager_GetNewsByFilters(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				repo: connDB,
			}
			list, err := m.GetNewsByFilters(context.Background(), tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetNewsByFilters() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.Len(t, list, tt.want, fmt.Sprint("len: ", len(list)))
		})
	}
}

func TestManager_GetNewsById(t *testing.T) {
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

			m := &Manager{
				repo: connDB,
			}

			news, err := m.GetNewsById(context.Background(), tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetById() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.NotNil(t, news, fmt.Sprint("no data found"))
			assert.Equal(t, news.Author, tt.want, fmt.Sprint("Author: ", news.Author))

		})
	}
}

func TestManager_GetNewsCount(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				repo: connDB,
			}

			count, err := m.GetNewsCount(context.Background(), tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("Count() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.Equal(t, tt.want, count, fmt.Sprint("Count() count: ", count))
		})
	}
}

func TestManager_AppNews(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *NewsInput
	}
	var timeSave = time.Now()
	tests := []struct {
		name    string
		args    args
		want    *News
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test 1 add full 'test 1'",
			args: args{
				ctx: context.Background(),
				in: &NewsInput{
					Title:      "test 1 title",
					Content:    nil,
					Author:     "test 1 author",
					CategoryID: 1,
					TagIDs:     []int{3, 2, 5},
				},
			},
			want: &News{
				News: db.News{
					Title:       "test 1 title",
					Content:     nil,
					Author:      "test 1 author",
					CategoryID:  1,
					TagIDs:      []int{3, 2, 5},
					PublishedAt: timeSave,
					StatusID:    1,
					Category:    &db.Category{ID: 1},
				},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				repo: connDB,
			}
			got, err := m.AddNews(tt.args.ctx, tt.args.in)
			if !tt.wantErr(t, err, fmt.Sprintf("AddNews(%v, %v)", tt.args.ctx, tt.args.in)) {
				return
			}
			assert.Equalf(t, tt.want, got, "AddNews(%v, %v)", tt.args.ctx, tt.args.in)
		})
	}
}
