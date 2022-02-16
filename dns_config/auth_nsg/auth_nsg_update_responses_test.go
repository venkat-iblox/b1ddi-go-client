package auth_nsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthNsgUpdateCreated_Error(t *testing.T) {
	resp := NewAuthNsgUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAuthNsgUpdateCreated_GetPayload(t *testing.T) {
	resp := NewAuthNsgUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
