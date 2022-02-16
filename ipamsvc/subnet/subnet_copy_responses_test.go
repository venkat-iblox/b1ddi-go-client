package subnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubnetCopyCreated_Error(t *testing.T) {
	resp := NewSubnetCopyCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestSubnetCopyCreated_GetPayload(t *testing.T) {
	resp := NewSubnetCopyCreated()
	assert.Nil(t, resp.GetPayload())
}
