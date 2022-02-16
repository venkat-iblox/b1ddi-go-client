package ipam_host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpamHostListOK_Error(t *testing.T) {
	resp := NewIpamHostListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestIpamHostListOK_GetPayload(t *testing.T) {
	resp := NewIpamHostListOK()
	assert.Nil(t, resp.GetPayload())
}
