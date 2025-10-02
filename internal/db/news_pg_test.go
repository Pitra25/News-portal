package db

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestNewsPG_GetAllByQuery(t *testing.T) {
	conn, err := getConnection()
	if err != nil {
		t.Log(err)
	}

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &NewsPG{
				db: conn,
			}
			list, err := m.GetAllByQuery(2, 1, 10, 1)
			if !tt.wantErr(t, err, fmt.Sprintf("GetAllByQuery() error = %e, wantErr %v", err, tt.wantErr)) {
				t.Log(err)
			}

			if assert.Len(t, list, tt.want) {
				t.Log(list)
			}
			assert.NotEqual(t, list, 0, "news in stock")
		})
	}
}

func TestNewsPG_GetById(t *testing.T) {
	conn, err := getConnection()
	if err != nil {
		t.Fatal(err)
	}

	m := &NewsPG{
		db: conn,
	}
	got, err := m.GetById(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(got)
}

func TestNewsPG_GetAllShortNews(t *testing.T) {
	conn, err := getConnection()
	if assert.NoError(t, err) {
		t.Fatal(err)
	}

	m := &NewsPG{
		db: conn,
	}
	lists, err := m.GetAllShortNews()
	if assert.NoError(t, err) {
		t.Fatal(err)
	}

	if assert.Len(t, lists, 0) {
		t.Log("no news")
	} else {
		t.Log(
			"len", len(lists),
			"element 1", lists[1],
		)
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
	if err != nil {
		t.Fatal(err)
	}

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

			assert.Equal(t, tt.want, count, "GetCount()", "count", count)
		})
	}
}
