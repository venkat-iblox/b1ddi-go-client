package forward_nsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardNsgCreateCreated_Error(t *testing.T) {
	resp := NewForwardNsgCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestForwardNsgCreateCreated_GetPayload(t *testing.T) {
	resp := NewForwardNsgCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
