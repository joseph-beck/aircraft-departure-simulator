package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewManageService(t *testing.T) {
	svc := NewManageService()
	assert.NotNil(t, svc)
}
