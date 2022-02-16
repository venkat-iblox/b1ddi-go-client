package option_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionSpaceUpdateCreated_Error(t *testing.T) {
	resp := NewOptionSpaceUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionSpaceUpdateCreated_GetPayload(t *testing.T) {
	resp := NewOptionSpaceUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
