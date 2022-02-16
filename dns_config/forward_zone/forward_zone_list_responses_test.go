package forward_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardZoneListOK_Error(t *testing.T) {
	resp := NewForwardZoneListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestForwardZoneListOK_GetPayload(t *testing.T) {
	resp := NewForwardZoneListOK()
	assert.Nil(t, resp.GetPayload())
}
