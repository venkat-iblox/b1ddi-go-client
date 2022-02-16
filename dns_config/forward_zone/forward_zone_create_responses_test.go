package forward_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardZoneCreateCreated_Error(t *testing.T) {
	resp := NewForwardZoneCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestForwardZoneCreateCreated_GetPayload(t *testing.T) {
	resp := NewForwardZoneCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
