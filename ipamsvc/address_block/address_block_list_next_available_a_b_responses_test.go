package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockListNextAvailableABOK_Error(t *testing.T) {
	resp := NewAddressBlockListNextAvailableABOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressBlockListNextAvailableABOK_GetPayload(t *testing.T) {
	resp := NewAddressBlockListNextAvailableABOK()
	assert.Nil(t, resp.GetPayload())
}
