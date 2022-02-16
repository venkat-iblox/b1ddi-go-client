package fixed_address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixedAddressUpdateCreated_Error(t *testing.T) {
	resp := NewFixedAddressUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestFixedAddressUpdateCreated_GetPayload(t *testing.T) {
	resp := NewFixedAddressUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
