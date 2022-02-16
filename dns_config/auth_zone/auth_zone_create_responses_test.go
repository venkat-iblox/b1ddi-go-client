package auth_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthZoneCreateCreated_Error(t *testing.T) {
	resp := NewAuthZoneCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestAuthZoneCreateCreated_GetPayload(t *testing.T) {
	resp := NewAuthZoneCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
