package auth_nsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthNsgListOK_Error(t *testing.T) {
	resp := NewAuthNsgListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAuthNsgListOK_GetPayload(t *testing.T) {
	resp := NewAuthNsgListOK()
	assert.Nil(t, resp.GetPayload())
}
