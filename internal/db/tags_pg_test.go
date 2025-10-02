package db

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestTagsPG_GetAll(t *testing.T) {
	conn, err := getConnection()
	assert.NoError(t, err)

	m := &TagsPG{
		db: conn,
	}
	tags, err := m.GetAll()
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(tags), minLength, fmt.Sprint("len: ", len(tags)))
}

func TestTagsPG_GetById(t *testing.T) {

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
		want    Tags
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get tags by id",
			fields: fields{
				db: conn,
			},
			args: args{
				id: 1,
			},
			want: Tags{
				TagID:    1,
				Title:    "Новости",
				StatusID: 1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "Get tags by id",
			fields: fields{
				db: conn,
			},
			args: args{
				id: 2,
			},
			want: Tags{
				TagID:    2,
				Title:    "Аналитика",
				StatusID: 1,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TagsPG{
				db: tt.fields.db,
			}
			got, err := m.GetById(tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("GetById(%d)", tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetById(%d)", tt.args.id)
		})
	}
}
