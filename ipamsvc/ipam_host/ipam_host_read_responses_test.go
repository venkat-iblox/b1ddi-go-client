package ipam_host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpamHostReadOK_Error(t *testing.T) {
	resp := NewIpamHostReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestIpamHostReadOK_GetPayload(t *testing.T) {
	resp := NewIpamHostReadOK()
	assert.Nil(t, resp.GetPayload())
}
