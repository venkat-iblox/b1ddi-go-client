package fixed_address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixedAddressCreateCreated_Error(t *testing.T) {
	resp := NewFixedAddressCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestFixedAddressCreateCreated_GetPayload(t *testing.T) {
	resp := NewFixedAddressCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
