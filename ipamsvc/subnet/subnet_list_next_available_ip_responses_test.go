package subnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubnetListNextAvailableIPOK_Error(t *testing.T) {
	resp := NewSubnetListNextAvailableIPOK()
	assert.NotEmpty(t, resp.Error())
}

func TestSubnetListNextAvailableIPOK_GetPayload(t *testing.T) {
	resp := NewSubnetListNextAvailableIPOK()
	assert.Nil(t, resp.GetPayload())
}
