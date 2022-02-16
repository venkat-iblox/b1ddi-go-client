package ip_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPSpaceCopyCreated_Error(t *testing.T) {
	resp := NewIPSpaceCopyCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestIPSpaceCopyCreated_GetPayload(t *testing.T) {
	resp := NewIPSpaceCopyCreated()
	assert.Nil(t, resp.GetPayload())
}
