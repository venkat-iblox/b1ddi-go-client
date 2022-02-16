package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServerCreateCreated_Error(t *testing.T) {
	resp := NewServerCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestServerCreateCreated_GetPayload(t *testing.T) {
	resp := NewServerCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
