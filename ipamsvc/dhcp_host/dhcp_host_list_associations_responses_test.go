package dhcp_host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDhcpHostListAssociationsOK_Error(t *testing.T) {
	resp := NewDhcpHostListAssociationsOK()
	assert.NotEmpty(t, resp.Error())
}

func TestDhcpHostListAssociationsOK_GetPayload(t *testing.T) {
	resp := NewDhcpHostListAssociationsOK()
	assert.Nil(t, resp.GetPayload())
}
