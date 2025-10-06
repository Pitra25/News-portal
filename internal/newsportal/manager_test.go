package newsportal

import (
	"News-portal/internal/db"
	"fmt"
	"log/slog"
	"os"
	"reflect"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var (
	connDB *db.DB
	opt    = pg.Options{
		Addr:     host + ":" + port,
		User:     user,
		Password: password,
		Database: dbname,
	}
)

func TestMain(m *testing.M) {
	conn, err := db.Connect(&opt)
	if err != nil {
		panic(err)
	}
	connDB = db.NewDB(conn)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestManager_GetALlShortNewsByFilters(t *testing.T) {
	type args struct {
		fil Filters
	}

	tests := []struct {
		name    string
		args    args
		want    []ShortNews
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				db: connDB,
			}
			got, err := m.GetALlShortNewsByFilters(tt.args.fil)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetALlShortNewsByFilters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetALlShortNewsByFilters() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_GetAllCategory(t *testing.T) {

	m := &Manager{
		db: connDB,
	}
	categories, err := m.GetAllCategory()
	assert.NoError(t, err)

	const minLength = 5
	slog.Info("categories", "cat: ", categories)
	assert.GreaterOrEqual(
		t, len(categories), minLength,
		fmt.Sprint("len: ", len(categories)),
	)
}

func TestManager_GetAllTag(t *testing.T) {

	m := &Manager{
		db: connDB,
	}
	tags, err := m.GetAllTag()
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(tags), minLength, fmt.Sprint("len: ", len(tags)))

}

func TestManager_GetNewsByFilters(t *testing.T) {

	tests := []struct {
		name    string
		args    Filters
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get all news by query (1)",
			args: Filters{
				NewsFilters{
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
			args: Filters{
				NewsFilters{
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
			args: Filters{
				NewsFilters{
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
			args: Filters{
				NewsFilters{
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
			m := &Manager{
				db: connDB,
			}
			list, err := m.GetNewsByFilters(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByFiltersNews() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			slog.Info("list", "len", len(list))

			assert.Len(t, list, tt.want, fmt.Sprint("len: ", len(list)))
		})
	}
}

func TestManager_GetNewsById(t *testing.T) {

	tests := []struct {
		name    string
		args    int
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Get news by id (1)",
			args:    1,
			want:    "Иван Петров",
			wantErr: assert.NoError,
		},
		{
			name:    "Get news by id (2)",
			args:    11,
			want:    "Анна Петрова",
			wantErr: assert.NoError,
		},
		{
			name:    "Get news by id (3)",
			args:    12,
			want:    "Михаил Семенов",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			m := &Manager{
				db: connDB,
			}

			news, err := m.GetNewsById(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetByFiltersNews() error = %e, wantErr %v", err, tt.wantErr)) {
				return
			}

			assert.NotNil(t, news, fmt.Sprint("no data found"))
			assert.Equal(t, news.Author, tt.want, fmt.Sprint("Author: ", news.Author))

		})
	}
}

func TestManager_GetNewsCount(t *testing.T) {

	tests := []struct {
		name    string
		args    Filters
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get count by categoryId 1 and tagId 0",
			args: Filters{
				News: NewsFilters{
					CategoryId: 1,
					TagId:      0,
				},
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 0",
			args: Filters{
				News: NewsFilters{
					CategoryId: 2,
					TagId:      0,
				},
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 1 and tagId 1",
			args: Filters{
				News: NewsFilters{
					CategoryId: 1,
					TagId:      1,
				},
			},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "Get count by categoryId 2 and tagId 1",
			args: Filters{
				News: NewsFilters{
					CategoryId: 2,
					TagId:      1,
				},
			},
			want:    1,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				db: connDB,
			}

			count, err := m.GetNewsCount(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("GetCount() error = %e, wantErr %v", err, tt.wantErr)) {

				return
			}

			assert.Equal(t, tt.want, count, fmt.Sprint("GetCount() count: ", count))
		})
	}
}
