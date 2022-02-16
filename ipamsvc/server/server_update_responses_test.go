package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServerUpdateCreated_Error(t *testing.T) {
	resp := NewServerUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestServerUpdateCreated_GetPayload(t *testing.T) {
	resp := NewServerUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
