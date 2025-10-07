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

func intPtr(i int) *int {
	return &i
}

func TestCategoryRepo_GetById(t *testing.T) {

	tests := []struct {
		name    string
		args    []int
		want    []Category
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "1",
			args: []int{1, 2, 3},
			want: []Category{
				{ID: 1, Title: "Политика", OrderNumber: intPtr(2), StatusID: 1},
				{ID: 2, Title: "Экономика", OrderNumber: intPtr(2), StatusID: 1},
				{ID: 3, Title: "Технологии", OrderNumber: intPtr(3), StatusID: 1},
			},
			wantErr: assert.NoError,
		},
		{
			name: "2",
			args: []int{2, 8, 5},
			want: []Category{
				{ID: 2, Title: "Экономика", OrderNumber: intPtr(2), StatusID: 1},
				{ID: 5, Title: "Культура", OrderNumber: intPtr(5), StatusID: 1},
				{ID: 8, Title: "Образование", OrderNumber: intPtr(8), StatusID: 1},
			},
			wantErr: assert.NoError,
		},
		{
			name: "3",
			args: []int{5, 1, 8},
			want: []Category{
				{ID: 1, Title: "Политика", OrderNumber: intPtr(1), StatusID: 1},
				{ID: 5, Title: "Культура", OrderNumber: intPtr(5), StatusID: 1},
				{ID: 8, Title: "Образование", OrderNumber: intPtr(8), StatusID: 1},
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
