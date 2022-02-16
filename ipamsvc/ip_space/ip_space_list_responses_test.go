package ip_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIPSpaceListOK_Error(t *testing.T) {
	resp := NewIPSpaceListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestIPSpaceListOK_GetPayload(t *testing.T) {
	resp := NewIPSpaceListOK()
	assert.Nil(t, resp.GetPayload())
}
