package ip_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPSpaceUpdateCreated_Error(t *testing.T) {
	resp := NewIPSpaceUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestIPSpaceUpdateCreated_GetPayload(t *testing.T) {
	resp := NewIPSpaceUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
