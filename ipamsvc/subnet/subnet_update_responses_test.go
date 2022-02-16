package subnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubnetUpdateCreated_Error(t *testing.T) {
	resp := NewSubnetUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestSubnetUpdateCreated_GetPayload(t *testing.T) {
	resp := NewSubnetUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
