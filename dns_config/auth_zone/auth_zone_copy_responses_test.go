package auth_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthZoneCopyCreated_Error(t *testing.T) {
	resp := NewAuthZoneCopyCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAuthZoneCopyCreated_GetPayload(t *testing.T) {
	resp := NewAuthZoneCopyCreated()
	assert.Nil(t, resp.GetPayload())
}
