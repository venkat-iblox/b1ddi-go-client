package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockListOK_Error(t *testing.T) {
	resp := NewAddressBlockListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressBlockListOK_GetPayload(t *testing.T) {
	resp := NewAddressBlockListOK()
	assert.Nil(t, resp.GetPayload())
}
