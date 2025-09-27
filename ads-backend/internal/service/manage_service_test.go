package service

import (
	"testing"

	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/internal/db"
	"github.com/stretchr/testify/assert"
)

type mockDB struct{}

func (m mockDB) Connect() (db.DB, error) {
	return db.DB{}, nil
}

func TestNewManageService(t *testing.T) {
	db := db.NewDB(mockDB{})
	svc := NewManageService(db)
	assert.NotNil(t, svc)
}
