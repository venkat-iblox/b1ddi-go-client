package global

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGlobalUpdateCreated_Error(t *testing.T) {
	resp := NewGlobalUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestGlobalUpdateCreated_GetPayload(t *testing.T) {
	resp := NewGlobalUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
