package global

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGlobalUpdate2Created_Error(t *testing.T) {
	resp := NewGlobalUpdate2Created()
	assert.NotEmpty(t, resp.Error())
}

func TestGlobalUpdate2Created_GetPayload(t *testing.T) {
	resp := NewGlobalUpdate2Created()
	assert.Nil(t, resp.GetPayload())
}
