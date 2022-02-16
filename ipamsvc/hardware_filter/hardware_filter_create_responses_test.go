package hardware_filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHardwareFilterCreateCreated_Error(t *testing.T) {
	resp := NewHardwareFilterCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestHardwareFilterCreateCreated_GetPayload(t *testing.T) {
	resp := NewHardwareFilterCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
