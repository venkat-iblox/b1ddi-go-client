package hardware_filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHardwareFilterListOK_Error(t *testing.T) {
	resp := NewHardwareFilterListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestHardwareFilterListOK_GetPayload(t *testing.T) {
	resp := NewHardwareFilterListOK()
	assert.Nil(t, resp.GetPayload())
}
