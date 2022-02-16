package ipam_host

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIpamHostDeleteNoContent_Error(t *testing.T) {
	resp := NewIpamHostDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
