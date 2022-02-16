package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServerListOK_Error(t *testing.T) {
	resp := NewServerListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestServerListOK_GetPayload(t *testing.T) {
	resp := NewServerListOK()
	assert.Nil(t, resp.GetPayload())
}
