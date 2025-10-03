package db

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestNewsPG_GetAllByQuery(t *testing.T) {
	conn, err := getConnection()
	assert.NoError(t, err)

	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		newsFilters FiltersNews
		pageFilters PageFilters
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get all news by query (1)",
			fields: fields{
				conn,
			},
			args: args{
				FiltersNews{
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
			fields: fields{
				conn,
			},
			args: args{
				FiltersNews{
					CategoryId: 4,
					TagId:      1,
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
			name: "Get all news by query (3)",
			fields: fields{
				conn,
			},
			args: args{
				FiltersNews{
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
			fields: fields{
				conn,
			},
			args: args{
				FiltersNews{
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &NewsRepo{
				db: conn,
			}
			list, err := m.GetByFilters(tt.args.newsFilters, tt.args.pageFilters)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByFiltersNews() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.Len(t, list, tt.want, fmt.Sprint("len: ", len(list)))
		})
	}
}

func TestNewsPG_GetById(t *testing.T) {
	conn, err := getConnection()
	assert.NoError(t, err)

	type fields struct {
		conn *sqlx.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    int
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get news by id (1)",
			fields: fields{
				conn,
			},
			args:    1,
			want:    "Иван Петров",
			wantErr: assert.NoError,
		},
		{
			name: "Get news by id (2)",
			fields: fields{
				conn,
			},
			args:    11,
			want:    "Анна Петрова",
			wantErr: assert.NoError,
		},
		{
			name: "Get news by id (3)",
			fields: fields{
				conn,
			},
			args:    12,
			want:    "Михаил Семенов",
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &NewsRepo{
				db: conn,
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
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		categoryId int
		tagId      int
	}

	conn, err := getConnection()
	assert.NoError(t, err)

	tests := []struct {
		name    string
		fields  fields
		args    FiltersNews
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get count by categoryId 1 and tagId 0",
			fields: fields{
				conn,
			},
			args: FiltersNews{
				CategoryId: 1,
				TagId:      0,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 0",
			fields: fields{
				conn,
			},
			args: FiltersNews{
				CategoryId: 2,
				TagId:      0,
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 1 and tagId 1",
			fields: fields{
				conn,
			},
			args: FiltersNews{
				CategoryId: 1,
				TagId:      1,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 1",
			fields: fields{
				conn,
			},
			args: FiltersNews{
				CategoryId: 2,
				TagId:      1,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 2",
			fields: fields{
				conn,
			},
			args: FiltersNews{
				CategoryId: 2,
				TagId:      2,
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 7 and tagId 1",
			fields: fields{
				conn,
			},
			args: FiltersNews{
				CategoryId: 7,
				TagId:      1,
			},
			want:    3,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &NewsRepo{
				db: tt.fields.db,
			}

			count, err := m.GetCount(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetCount() error = %e, wantErr %v", err, tt.wantErr)) {

				return
			}

			assert.Equal(t, tt.want, count, fmt.Sprint("GetCount() count: ", count))
		})
	}
}
