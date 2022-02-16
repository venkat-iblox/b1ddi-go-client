package dhcp_host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDhcpHostListOK_Error(t *testing.T) {
	resp := NewDhcpHostListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestDhcpHostListOK_GetPayload(t *testing.T) {
	resp := NewDhcpHostListOK()
	assert.Nil(t, resp.GetPayload())
}
