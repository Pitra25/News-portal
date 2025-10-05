package db

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoriesPG_GetAll(t *testing.T) {

	m := &CategoryRepo{
		db: connDB,
	}

	categories, err := m.GetAll()
	assert.NoError(t, err)

	const minLength = 5
	slog.Info("categories", "cat: ", categories)
	assert.GreaterOrEqual(
		t, len(categories), minLength,
		fmt.Sprint("len: ", len(categories)),
	)
}

func TestCategoryRepo_GetById(t *testing.T) {

	tests := []struct {
		name    string
		args    []int
		want    []Categories
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "1",
			args: []int{1, 2, 3},
			want: []Categories{
				{CategoryID: 1, Title: "Политика", OrderNumber: 1, StatusID: 1},
				{CategoryID: 2, Title: "Экономика", OrderNumber: 2, StatusID: 1},
				{CategoryID: 3, Title: "Технологии", OrderNumber: 3, StatusID: 1},
			},
			wantErr: assert.NoError,
		},
		{
			name: "2",
			args: []int{2, 8, 5},
			want: []Categories{
				{CategoryID: 2, Title: "Экономика", OrderNumber: 2, StatusID: 1},
				{CategoryID: 5, Title: "Культура", OrderNumber: 5, StatusID: 1},
				{CategoryID: 8, Title: "Образование", OrderNumber: 8, StatusID: 1},
			},
			wantErr: assert.NoError,
		},
		{
			name: "3",
			args: []int{5, 1, 8},
			want: []Categories{
				{CategoryID: 1, Title: "Политика", OrderNumber: 1, StatusID: 1},
				{CategoryID: 5, Title: "Культура", OrderNumber: 5, StatusID: 1},
				{CategoryID: 8, Title: "Образование", OrderNumber: 8, StatusID: 1},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &CategoryRepo{
				db: connDB,
			}
			got, err := m.GetById(tt.args)
			if !tt.wantErr(t, err,
				fmt.Sprintf("GetById(%v)", tt.args)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetById(%v)", tt.args)
		})
	}
}
