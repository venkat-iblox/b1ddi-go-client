package subnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubnetCreateNextAvailableIPCreated_Error(t *testing.T) {
	resp := NewSubnetCreateNextAvailableIPCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestSubnetCreateNextAvailableIPCreated_GetPayload(t *testing.T) {
	resp := NewSubnetCreateNextAvailableIPCreated()
	assert.Nil(t, resp.GetPayload())
}
