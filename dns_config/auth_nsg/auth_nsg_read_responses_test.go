package auth_nsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthNsgReadOK_Error(t *testing.T) {
	resp := NewAuthNsgReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAuthNsgReadOK_GetPayload(t *testing.T) {
	resp := NewAuthNsgReadOK()
	assert.Nil(t, resp.GetPayload())
}
