package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagsPG_GetAll(t *testing.T) {
	conn, err := getConnection()
	assert.NoError(t, err)

	m := &TagRepo{
		db: conn,
	}
	tags, err := m.GetAll()
	assert.NoError(t, err)

	const minLength = 5
	assert.GreaterOrEqual(t, len(tags), minLength, fmt.Sprint("len: ", len(tags)))
}
