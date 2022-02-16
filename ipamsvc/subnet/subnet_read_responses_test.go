package subnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubnetReadOK_Error(t *testing.T) {
	resp := NewSubnetReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestSubnetReadOK_GetPayload(t *testing.T) {
	resp := NewSubnetReadOK()
	assert.Nil(t, resp.GetPayload())
}
