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

func TestManageServicePing(t *testing.T) {
	db := db.NewDB(mockDB{})
	svc := NewManageService(db)

	var reply string
	err := svc.Ping("test", &reply)
	assert.NoError(t, err)
	assert.Equal(t, "ping: test", reply)
}
