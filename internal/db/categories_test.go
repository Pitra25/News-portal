package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoriesPG_GetAll(t *testing.T) {
	m := &CategoryRepo{
		db: connDB,
	}
	tags, err := m.GetAll()
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(tags), minLength, fmt.Sprint("len: ", len(tags)))
}

func TestCategoryRepo_GetById(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		want    Categories
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "1",
			args:    1,
			want:    Categories{CategoryID: 1, Title: "Политика", OrderNumber: 1, StatusID: 1},
			wantErr: assert.NoError,
		},
		{
			name:    "2",
			args:    2,
			want:    Categories{CategoryID: 2, Title: "Экономика", OrderNumber: 2, StatusID: 1},
			wantErr: assert.NoError,
		},
		{
			name:    "3",
			args:    3,
			want:    Categories{CategoryID: 3, Title: "Технологии", OrderNumber: 3, StatusID: 1},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &CategoryRepo{
				db: connDB,
			}
			got, err := m.GetById(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetById(%v)", tt.args)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetById(%v)", tt.args)
		})
	}
}
