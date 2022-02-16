package hardware_filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHardwareFilterReadOK_Error(t *testing.T) {
	resp := NewHardwareFilterReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestHardwareFilterReadOK_GetPayload(t *testing.T) {
	resp := NewHardwareFilterReadOK()
	assert.Nil(t, resp.GetPayload())
}
