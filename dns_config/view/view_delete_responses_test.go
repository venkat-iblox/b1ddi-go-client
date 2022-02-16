package view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewDeleteNoContent_Error(t *testing.T) {
	resp := NewViewDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
