package address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressReadOK_Error(t *testing.T) {
	resp := NewAddressReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressReadOK_GetPayload(t *testing.T) {
	resp := NewAddressReadOK()
	assert.Nil(t, resp.GetPayload())
}
