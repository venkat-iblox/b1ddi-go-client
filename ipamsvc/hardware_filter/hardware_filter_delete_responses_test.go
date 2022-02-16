package hardware_filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHardwareFilterDeleteNoContent_Error(t *testing.T) {
	resp := NewHardwareFilterDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
