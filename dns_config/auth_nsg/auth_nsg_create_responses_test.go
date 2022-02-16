package auth_nsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthNsgCreateCreated_Error(t *testing.T) {
	resp := NewAuthNsgCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAuthNsgCreateCreated_GetPayload(t *testing.T) {
	resp := NewAuthNsgCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
