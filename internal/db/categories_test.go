package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoriesPG_GetAll(t *testing.T) {
	conn, err := getConnection()
	assert.NoError(t, err)

	m := &CategoryRepo{
		db: conn,
	}
	tags, err := m.GetAll()
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(tags), minLength, fmt.Sprint("len: ", len(tags)))
}
