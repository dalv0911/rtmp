package rtmp_test

import (
	"testing"

	"github.com/dalv0911/rtmp"
	"github.com/dalv0911/rtmp/server"
	"github.com/stretchr/testify/assert"
)

func TestNewServerMakesNewServers(t *testing.T) {
	s, err := rtmp.NewServer(":1935")

	assert.IsType(t, new(server.Server), s)
	assert.Nil(t, err)
}
