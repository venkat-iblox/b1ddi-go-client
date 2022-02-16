package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockCreateCreated_Error(t *testing.T) {
	resp := NewAddressBlockCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressBlockCreateCreated_GetPayload(t *testing.T) {
	resp := NewAddressBlockCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
