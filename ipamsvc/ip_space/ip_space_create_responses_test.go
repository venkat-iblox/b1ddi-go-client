package ip_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPSpaceCreateCreated_Error(t *testing.T) {
	resp := NewIPSpaceCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestIPSpaceCreateCreated_GetPayload(t *testing.T) {
	resp := NewIPSpaceCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
