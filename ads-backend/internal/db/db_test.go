package db

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockDB struct{}

func (m mockDB) Connect() (DB, error) {
	return DB{}, nil
}

type mockErr struct{}

func (m mockErr) Connect() (DB, error) {
	return DB{}, errors.New("error")
}

func TestNewDB(t *testing.T) {
	connDB := mockDB{}

	assert.NotPanics(t, func() { NewDB(connDB) })

	db := NewDB(connDB)
	assert.NotNil(t, db)

	connErr := mockErr{}
	assert.Panics(t, func() { NewDB(connErr) })
}
