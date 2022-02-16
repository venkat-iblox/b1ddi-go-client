package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockUpdateCreated_Error(t *testing.T) {
	resp := NewAddressBlockUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressBlockUpdateCreated_GetPayload(t *testing.T) {
	resp := NewAddressBlockUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
