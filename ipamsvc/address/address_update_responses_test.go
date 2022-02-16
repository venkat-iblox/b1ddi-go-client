package address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressUpdateCreated_Error(t *testing.T) {
	resp := NewAddressUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressUpdateCreated_GetPayload(t *testing.T) {
	resp := NewAddressUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
