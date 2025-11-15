package adsmanage

import (
	"net/rpc"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	client, err := rpc.DialHTTP("tcp", "localhost:4040")
	assert.NoError(t, err)

	defer client.Close()

	var resp string
	err = client.Call("ManageService.Ping", "pong", &resp)
	assert.NoError(t, err)
	assert.Equal(t, "ping: pong", resp)

	err = client.Call("ManageService.Ping", "ping", &resp)
	assert.NoError(t, err)
	assert.Equal(t, "ping: ping", resp)
}
