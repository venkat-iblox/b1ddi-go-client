package fixed_address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixedAddressReadOK_Error(t *testing.T) {
	resp := NewFixedAddressReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestFixedAddressReadOK_GetPayload(t *testing.T) {
	resp := NewFixedAddressReadOK()
	assert.Nil(t, resp.GetPayload())
}
