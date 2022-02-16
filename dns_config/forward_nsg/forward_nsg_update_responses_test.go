package forward_nsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardNsgUpdateCreated_Error(t *testing.T) {
	resp := NewForwardNsgUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestForwardNsgUpdateCreated_GetPayload(t *testing.T) {
	resp := NewForwardNsgUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
