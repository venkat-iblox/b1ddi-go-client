package address

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressCreateCreated_Error(t *testing.T) {
	resp := NewAddressCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAddressCreateCreated_GetPayload(t *testing.T) {
	resp := NewAddressCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
