package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServerReadOK_Error(t *testing.T) {
	resp := NewServerReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestServerReadOK_GetPayload(t *testing.T) {
	resp := NewServerReadOK()
	assert.Nil(t, resp.GetPayload())
}
