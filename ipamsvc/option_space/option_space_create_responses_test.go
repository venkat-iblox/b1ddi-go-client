package option_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionSpaceCreateCreated_Error(t *testing.T) {
	resp := NewOptionSpaceCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionSpaceCreateCreated_GetPayload(t *testing.T) {
	resp := NewOptionSpaceCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
