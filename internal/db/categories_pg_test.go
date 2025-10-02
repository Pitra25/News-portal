package db

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCategoriesPG_GetAll(t *testing.T) {
	conn, err := getConnection()
	assert.NoError(t, err)

	m := &CategoriesPG{
		db: conn,
	}
	tags, err := m.GetAll()
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(tags), minLength, fmt.Sprint("len: ", len(tags)))
}

func TestCategoriesPG_GetById(t *testing.T) {

	conn, err := getConnection()
	assert.NoError(t, err)

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Categories
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success",
			fields: fields{
				db: conn,
			},
			args: args{
				id: 1,
			},
			want: Categories{
				CategoryID:  1,
				Title:       "Политика",
				OrderNumber: 1,
				StatusID:    1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success",
			fields: fields{
				db: conn,
			},
			args: args{
				id: 2,
			},
			want: Categories{
				CategoryID:  2,
				Title:       "Экономика",
				OrderNumber: 2,
				StatusID:    1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success",
			fields: fields{
				db: conn,
			},
			args: args{
				id: 3,
			},
			want: Categories{
				CategoryID:  3,
				Title:       "Технологии",
				OrderNumber: 3,
				StatusID:    1,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &CategoriesPG{
				db: tt.fields.db,
			}
			got, err := m.GetById(tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("GetById(%v)", tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetById(%v)", tt.args.id)
		})
	}
}
