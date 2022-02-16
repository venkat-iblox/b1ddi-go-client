package auth_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthZoneDeleteNoContent_Error(t *testing.T) {
	resp := NewAuthZoneDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
