package auth_nsg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthNsgDeleteNoContent_Error(t *testing.T) {
	resp := NewAuthNsgDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
