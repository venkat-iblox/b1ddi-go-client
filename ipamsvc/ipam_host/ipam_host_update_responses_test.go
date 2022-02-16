package ipam_host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpamHostUpdateCreated_Error(t *testing.T) {
	resp := NewIpamHostUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestIpamHostUpdateCreated_GetPayload(t *testing.T) {
	resp := NewIpamHostUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
