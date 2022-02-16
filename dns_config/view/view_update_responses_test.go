package view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewUpdateCreated_Error(t *testing.T) {
	resp := NewViewUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestViewUpdateCreated_GetPayload(t *testing.T) {
	resp := NewViewUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
