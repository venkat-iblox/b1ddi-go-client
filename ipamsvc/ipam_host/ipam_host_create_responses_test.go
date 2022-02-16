package ipam_host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpamHostCreateCreated_Error(t *testing.T) {
	resp := NewIpamHostCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestIpamHostCreateCreated_GetPayload(t *testing.T) {
	resp := NewIpamHostCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
