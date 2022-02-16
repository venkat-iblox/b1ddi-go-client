package subnet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubnetListOK_Error(t *testing.T) {
	resp := NewSubnetListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestSubnetListOK_GetPayload(t *testing.T) {
	resp := NewSubnetListOK()
	assert.Nil(t, resp.GetPayload())
}
