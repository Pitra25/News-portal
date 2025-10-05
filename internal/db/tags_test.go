package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsPG_GetAll(t *testing.T) {
	m := &TagRepo{
		db: connDB,
	}
	tags, err := m.GetAll()
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(tags), minLength, fmt.Sprint("len: ", len(tags)))
}

func TestTagRepo_GetByID(t *testing.T) {

	tests := []struct {
		name    string
		args    []int
		want    []Tags
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "test 1",
			args: []int{1, 2, 10},
			want: []Tags{
				{TagID: 1, Title: "Новости", StatusID: 1},
				{TagID: 2, Title: "Аналитика", StatusID: 1},
				{TagID: 10, Title: "Прогноз", StatusID: 1},
			},
			wantErr: assert.NoError,
		},
		{
			name: "test 2",
			args: []int{1, 6, 9},
			want: []Tags{
				{TagID: 1, Title: "Новости", StatusID: 1},
				{TagID: 6, Title: "Обзор", StatusID: 1},
				{TagID: 9, Title: "Статистика", StatusID: 1},
			},
			wantErr: assert.NoError,
		},
		{
			name: "test 3",
			args: []int{1, 5, 9},
			want: []Tags{
				{TagID: 1, Title: "Новости", StatusID: 1},
				{TagID: 5, Title: "Реportаж", StatusID: 1},
				{TagID: 9, Title: "Статистика", StatusID: 1},
			},
			wantErr: assert.NoError,
		},
		{
			name: "test 4",
			args: []int{1, 4, 5},
			want: []Tags{
				{TagID: 1, Title: "Новости", StatusID: 1},
				{TagID: 4, Title: "Интервью", StatusID: 1},
				{TagID: 5, Title: "Реportаж", StatusID: 1},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TagRepo{
				db: connDB,
			}
			got, err := m.GetByID(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByID(%v)", tt.args)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetByID(%v)", tt.args)
		})
	}
}
