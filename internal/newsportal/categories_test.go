package newsportal

import (
	"News-portal/internal/db"
	"reflect"
	"testing"
)

func TestCategoriesService_GetAll(t *testing.T) {
	type fields struct {
		db *db.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Categories
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CategoriesService{
				db: tt.fields.db,
			}
			got, err := s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCategoriesService_GetById(t *testing.T) {
	type fields struct {
		db *db.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Categories
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CategoriesService{
				db: tt.fields.db,
			}
			got, err := s.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
