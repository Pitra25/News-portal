package db

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// не работает, пока что
func TestNewsPG_GetAll(t *testing.T) {
	conn, err := getConnection()
	assert.NoError(t, err)

	m := &NewsPG{
		db: conn,
	}
	// TODO изменить sql запрос
	lists, err := m.GetAll()
	assert.NoError(t, err)

	const minLength = 10
	assert.GreaterOrEqual(t, len(lists), minLength, fmt.Sprint("len: ", len(lists)))
}

func TestNewsPG_GetAllByQuery(t *testing.T) {
	conn, err := getConnection()
	assert.NoError(t, err)

	type fields struct {
		conn *sqlx.DB
	}
	type args struct {
		categoryId int
		tagId      int
		pageSize   int
		page       int
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
				categoryId: 2,
				tagId:      1,
				pageSize:   10,
				page:       1,
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
				categoryId: 4,
				tagId:      1,
				pageSize:   10,
				page:       1,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (3)",
			fields: fields{
				conn,
			},
			args: args{
				categoryId: 5,
				tagId:      1,
				pageSize:   10,
				page:       1,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "Get all news by query (4)",
			fields: fields{
				conn,
			},
			args: args{
				categoryId: 10,
				tagId:      1,
				pageSize:   10,
				page:       1,
			},
			want:    1,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &NewsPG{
				db: conn,
			}
			list, err := m.GetAllByQuery(2, 1, 10, 1)
			if !tt.wantErr(t, err, fmt.Sprintf("GetAllByQuery() error = %e, wantErr %v", err, tt.wantErr)) {
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
	type args struct {
		newsId int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get news by id (1)",
			fields: fields{
				conn,
			},
			args: args{
				newsId: 1,
			},
			want:    "Иван Петров",
			wantErr: assert.NoError,
		},
		{
			name: "Get news by id (2)",
			fields: fields{
				conn,
			},
			args: args{
				newsId: 11,
			},
			want:    "Анна Петрова",
			wantErr: assert.NoError,
		},
		{
			name: "Get news by id (3)",
			fields: fields{
				conn,
			},
			args: args{
				newsId: 12,
			},
			want:    "Михаил Семенов",
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &NewsPG{
				db: conn,
			}
			news, err := m.GetById(tt.args.newsId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetAllByQuery() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.NotNil(t, news, fmt.Sprint("no data found"))
			assert.Equal(t, news.Author, tt.want, fmt.Sprint("Author: ", news.Author))
		})
	}
}

func TestNewsPG_GetAllShortNews(t *testing.T) {
	conn, err := getConnection()
	assert.NoError(t, err)

	m := &NewsPG{
		db: conn,
	}
	lists, err := m.GetAllShortNews()
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(lists), minLength, fmt.Sprint("len: ", len(lists)))
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
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get count by categoryId 1 and tagId 0",
			fields: fields{
				conn,
			},
			args: args{
				categoryId: 1,
				tagId:      0,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 0",
			fields: fields{
				conn,
			},
			args: args{
				categoryId: 2,
				tagId:      0,
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 1 and tagId 1",
			fields: fields{
				conn,
			},
			args: args{
				categoryId: 1,
				tagId:      1,
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 1",
			fields: fields{
				conn,
			},
			args: args{
				categoryId: 2,
				tagId:      1,
			},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 2",
			fields: fields{
				conn,
			},
			args: args{
				categoryId: 2,
				tagId:      2,
			},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 7 and tagId 1",
			fields: fields{
				conn,
			},
			args: args{
				categoryId: 7,
				tagId:      1,
			},
			want:    3,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &NewsPG{
				db: tt.fields.db,
			}

			count, err := m.GetCount(tt.args.categoryId, tt.args.tagId)
			if !tt.wantErr(t, err, fmt.Sprintf("GetCount() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.Equal(t, tt.want, count, fmt.Sprint("GetCount() count: ", count))
		})
	}
}
