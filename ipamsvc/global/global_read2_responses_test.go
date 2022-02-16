package global

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGlobalRead2OK_Error(t *testing.T) {
	resp := NewGlobalRead2OK()
	assert.NotEmpty(t, resp.Error())
}

func TestGlobalRead2OK_GetPayload(t *testing.T) {
	resp := NewGlobalRead2OK()
	assert.Nil(t, resp.GetPayload())
}
