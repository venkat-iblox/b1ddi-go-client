package forward_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardZoneReadOK_Error(t *testing.T) {
	resp := NewForwardZoneReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestForwardZoneReadOK_GetPayload(t *testing.T) {
	resp := NewForwardZoneReadOK()
	assert.Nil(t, resp.GetPayload())
}
