package forward_nsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardNsgListOK_Error(t *testing.T) {
	resp := NewForwardNsgListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestForwardNsgListOK_GetPayload(t *testing.T) {
	resp := NewForwardNsgListOK()
	assert.Nil(t, resp.GetPayload())
}
