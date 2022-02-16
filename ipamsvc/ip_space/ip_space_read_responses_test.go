package ip_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPSpaceReadOK_Error(t *testing.T) {
	resp := NewIPSpaceReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestIPSpaceReadOK_GetPayload(t *testing.T) {
	resp := NewIPSpaceReadOK()
	assert.Nil(t, resp.GetPayload())
}
