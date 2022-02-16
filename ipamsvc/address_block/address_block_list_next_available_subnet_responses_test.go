package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockListNextAvailableSubnetOK_Error(t *testing.T) {
	resp := NewAddressBlockListNextAvailableSubnetOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressBlockListNextAvailableSubnetOK_GetPayload(t *testing.T) {
	resp := NewAddressBlockListNextAvailableSubnetOK()
	assert.Nil(t, resp.GetPayload())
}
