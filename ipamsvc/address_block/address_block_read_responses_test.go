package address_block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressBlockReadOK_Error(t *testing.T) {
	resp := NewAddressBlockReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressBlockReadOK_GetPayload(t *testing.T) {
	resp := NewAddressBlockReadOK()
	assert.Nil(t, resp.GetPayload())
}
