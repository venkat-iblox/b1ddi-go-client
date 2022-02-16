package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockCreateNextAvailableABCreated_Error(t *testing.T) {
	resp := NewAddressBlockCreateNextAvailableABCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressBlockCreateNextAvailableABCreated_GetPayload(t *testing.T) {
	resp := NewAddressBlockCreateNextAvailableABCreated()
	assert.Nil(t, resp.GetPayload())
}
