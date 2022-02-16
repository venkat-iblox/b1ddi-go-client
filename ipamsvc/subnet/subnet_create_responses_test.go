package subnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubnetCreateCreated_Error(t *testing.T) {
	resp := NewSubnetCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestSubnetCreateCreated_GetPayload(t *testing.T) {
	resp := NewSubnetCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
