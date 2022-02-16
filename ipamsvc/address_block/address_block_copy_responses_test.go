package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockCopyCreated_Error(t *testing.T) {
	resp := NewAddressBlockCopyCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressBlockCopyCreated_GetPayload(t *testing.T) {
	resp := NewAddressBlockCopyCreated()
	assert.Nil(t, resp.GetPayload())
}
