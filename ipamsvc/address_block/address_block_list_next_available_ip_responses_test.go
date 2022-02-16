package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockListNextAvailableIPOK_Error(t *testing.T) {
	resp := NewAddressBlockListNextAvailableIPOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressBlockListNextAvailableIPOK_GetPayload(t *testing.T) {
	resp := NewAddressBlockListNextAvailableIPOK()
	assert.Nil(t, resp.GetPayload())
}
