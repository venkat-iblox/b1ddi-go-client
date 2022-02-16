package hardware_filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHardwareFilterUpdateCreated_Error(t *testing.T) {
	resp := NewHardwareFilterUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestHardwareFilterUpdateCreated_GetPayload(t *testing.T) {
	resp := NewHardwareFilterUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
