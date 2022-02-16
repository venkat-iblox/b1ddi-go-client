package ha_group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHaGroupReadOK_Error(t *testing.T) {
	resp := NewHaGroupReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestHaGroupReadOK_GetPayload(t *testing.T) {
	resp := NewHaGroupReadOK()
	assert.Nil(t, resp.GetPayload())
}
