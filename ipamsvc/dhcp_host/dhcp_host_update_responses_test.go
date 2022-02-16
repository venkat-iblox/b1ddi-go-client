package dhcp_host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDhcpHostUpdateCreated_Error(t *testing.T) {
	resp := NewDhcpHostUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestDhcpHostUpdateCreated_GetPayload(t *testing.T) {
	resp := NewDhcpHostUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
