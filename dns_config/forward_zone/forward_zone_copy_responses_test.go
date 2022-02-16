package forward_zone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForwardZoneCopyCreated_Error(t *testing.T) {
	resp := NewForwardZoneCopyCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestForwardZoneCopyCreated_GetPayload(t *testing.T) {
	resp := NewForwardZoneCopyCreated()
	assert.Nil(t, resp.GetPayload())
}
