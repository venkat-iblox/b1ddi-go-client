package dhcp_host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDhcpHostReadOK_Error(t *testing.T) {
	resp := NewDhcpHostReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestDhcpHostReadOK_GetPayload(t *testing.T) {
	resp := NewDhcpHostReadOK()
	assert.Nil(t, resp.GetPayload())
}
