package fixed_address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixedAddressListOK_Error(t *testing.T) {
	resp := NewFixedAddressListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestFixedAddressListOK_GetPayload(t *testing.T) {
	resp := NewFixedAddressListOK()
	assert.Nil(t, resp.GetPayload())
}
