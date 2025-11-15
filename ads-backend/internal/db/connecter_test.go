package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCredentials(t *testing.T) {
	creds := NewCredentials("localhost", "5432", "user", "password")

	assert.Equal(t, "localhost", creds.Host)
	assert.Equal(t, "5432", creds.Port)
	assert.Equal(t, "user", creds.User)
	assert.Equal(t, "password", creds.Password)
}

func TestDefaultCredentials(t *testing.T) {
	t.Setenv("PG_HOST", "localhost")
	t.Setenv("PG_PORT", "5432")
	t.Setenv("PG_USER", "user")
	t.Setenv("PG_PASS", "password")

	creds := DefaultCredentials()

	assert.Equal(t, "localhost", creds.Host)
	assert.Equal(t, "5432", creds.Port)
	assert.Equal(t, "user", creds.User)
	assert.Equal(t, "password", creds.Password)
}

func TestPGConnecterURL(t *testing.T) {
	creds := NewCredentials("localhost", "5432", "user", "password")
	conn := NewPGConnecter("test", creds)

	url := conn.url()
	expected := "postgres://user:password@localhost:5432"

	assert.Equal(t, expected, url)
}
