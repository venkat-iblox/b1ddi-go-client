package view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewCreateCreated_Error(t *testing.T) {
	resp := NewViewCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestViewCreateCreated_GetPayload(t *testing.T) {
	resp := NewViewCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
