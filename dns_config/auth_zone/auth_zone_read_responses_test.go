package auth_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthZoneReadOK_Error(t *testing.T) {
	resp := NewAuthZoneReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAuthZoneReadOK_GetPayload(t *testing.T) {
	resp := NewAuthZoneReadOK()
	assert.Nil(t, resp.GetPayload())
}
