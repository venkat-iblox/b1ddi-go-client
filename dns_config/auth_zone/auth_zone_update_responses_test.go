package auth_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthZoneUpdateCreated_Error(t *testing.T) {
	resp := NewAuthZoneUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAuthZoneUpdateCreated_GetPayload(t *testing.T) {
	resp := NewAuthZoneUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
