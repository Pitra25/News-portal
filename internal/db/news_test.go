package db

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewsPG_GetAllByQuery(t *testing.T) {
	tests := []struct {
		name    string
		args    Filters
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get all news by query (1)",
			args: Filters{
				NewsFilters{
					CategoryId: 2,
					TagId:      1,
				},
				PageFilters{
					PageSize: 10,
					Page:     1,
				},
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (2)",
			args: Filters{
				NewsFilters{
					CategoryId: 4,
					TagId:      1,
				},
				PageFilters{
					PageSize: 10,
					Page:     1,
				},
			},
			want:    4,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (3)",
			args: Filters{
				NewsFilters{
					CategoryId: 5,
					TagId:      1,
				},
				PageFilters{
					PageSize: 10,
					Page:     1,
				},
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (4)",
			args: Filters{
				NewsFilters{
					CategoryId: 10,
					TagId:      1,
				},
				PageFilters{
					PageSize: 10,
					Page:     1,
				},
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (4)",
			args: Filters{
				NewsFilters{
					CategoryId: 10,
					TagId:      0,
				},
				PageFilters{
					PageSize: 10,
					Page:     1,
				},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (4)",
			args: Filters{
				NewsFilters{
					CategoryId: 0,
					TagId:      1,
				},
				PageFilters{
					PageSize: 10,
					Page:     1,
				},
			},
			want:    10,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &NewsRepo{
				db: connDB,
			}
			list, err := m.GetByFilters(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByFiltersNews() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			slog.Info("list", "len", len(list))

			assert.Len(t, list, tt.want, fmt.Sprint("len: ", len(list)))
		})
	}
}

func TestNewsPG_GetById(t *testing.T) {
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

			m := &NewsRepo{
				db: connDB,
			}
			news, err := m.GetById(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByFiltersNews() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.NotNil(t, news, fmt.Sprint("no data found"))
			assert.Equal(t, news.Author, tt.want, fmt.Sprint("Author: ", news.Author))
		})
	}
}

func TestNewsPG_GetCount(t *testing.T) {
	tests := []struct {
		name    string
		args    Filters
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get count by categoryId 1 and tagId 0",
			args: Filters{
				News: NewsFilters{
					CategoryId: 1,
					TagId:      0,
				},
			},
			want:    3,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 0",
			args: Filters{
				News: NewsFilters{
					CategoryId: 2,
					TagId:      0,
				},
			},
			want:    3,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 1 and tagId 1",
			args: Filters{
				News: NewsFilters{
					CategoryId: 1,
					TagId:      1,
				},
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 1",
			args: Filters{
				News: NewsFilters{
					CategoryId: 2,
					TagId:      1,
				},
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 2",
			args: Filters{
				News: NewsFilters{
					CategoryId: 2,
					TagId:      2,
				},
			},
			want:    3,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 7 and tagId 1",
			args: Filters{
				News: NewsFilters{
					CategoryId: 7,
					TagId:      1,
				},
			},
			want:    4,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &NewsRepo{
				db: connDB,
			}

			count, err := m.GetCount(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetCount() error = %e, wantErr %v", err, tt.wantErr)) {

				return
			}

			assert.Equal(t, tt.want, count, fmt.Sprint("GetCount() count: ", count))
		})
	}
}
