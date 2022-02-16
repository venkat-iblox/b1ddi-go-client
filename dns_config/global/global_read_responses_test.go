package global

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGlobalReadOK_Error(t *testing.T) {
	resp := NewGlobalReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestGlobalReadOK_GetPayload(t *testing.T) {
	resp := NewGlobalReadOK()
	assert.Nil(t, resp.GetPayload())
}
