package forward_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardZoneDeleteNoContent_Error(t *testing.T) {
	resp := NewForwardZoneDeleteNoContent()
	assert.NotEmpty(t, resp.Error())
}
