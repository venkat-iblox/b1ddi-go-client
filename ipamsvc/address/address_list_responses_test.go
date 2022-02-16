package address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressListOK_Error(t *testing.T) {
	resp := NewAddressListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressListOK_GetPayload(t *testing.T) {
	resp := NewAddressListOK()
	assert.Nil(t, resp.GetPayload())
}
