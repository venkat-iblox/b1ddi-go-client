package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockCreateNextAvailableSubnetCreated_Error(t *testing.T) {
	resp := NewAddressBlockCreateNextAvailableSubnetCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressBlockCreateNextAvailableSubnetCreated_GetPayload(t *testing.T) {
	resp := NewAddressBlockCreateNextAvailableSubnetCreated()
	assert.Nil(t, resp.GetPayload())
}
