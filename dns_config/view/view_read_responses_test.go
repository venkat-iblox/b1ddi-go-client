package view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewReadOK_Error(t *testing.T) {
	resp := NewViewReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestViewReadOK_GetPayload(t *testing.T) {
	resp := NewViewReadOK()
	assert.Nil(t, resp.GetPayload())
}
