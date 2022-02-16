package view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewBulkCopyCreated_Error(t *testing.T) {
	resp := NewViewBulkCopyCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestViewBulkCopyCreated_GetPayload(t *testing.T) {
	resp := NewViewBulkCopyCreated()
	assert.Nil(t, resp.GetPayload())
}
