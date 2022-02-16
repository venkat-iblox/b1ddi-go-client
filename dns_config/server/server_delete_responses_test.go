package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServerDeleteNoContent_Error(t *testing.T) {
	resp := NewServerDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
