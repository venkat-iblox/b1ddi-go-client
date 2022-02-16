package forward_nsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardNsgReadOK_Error(t *testing.T) {
	resp := NewForwardNsgReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestForwardNsgReadOK_GetPayload(t *testing.T) {
	resp := NewForwardNsgReadOK()
	assert.Nil(t, resp.GetPayload())
}
