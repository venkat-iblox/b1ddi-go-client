package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockCreateNextAvailableIPCreated_Error(t *testing.T) {
	resp := NewAddressBlockCreateNextAvailableIPCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressBlockCreateNextAvailableIPCreated_GetPayload(t *testing.T) {
	resp := NewAddressBlockCreateNextAvailableIPCreated()
	assert.Nil(t, resp.GetPayload())
}
