package forward_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardZoneUpdateCreated_Error(t *testing.T) {
	resp := NewForwardZoneUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestForwardZoneUpdateCreated_GetPayload(t *testing.T) {
	resp := NewForwardZoneUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
