package auth_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthZoneListOK_Error(t *testing.T) {
	resp := NewAuthZoneListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestAuthZoneListOK_GetPayload(t *testing.T) {
	resp := NewAuthZoneListOK()
	assert.Nil(t, resp.GetPayload())
}
