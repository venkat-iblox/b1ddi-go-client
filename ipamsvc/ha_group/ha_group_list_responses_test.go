package ha_group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHaGroupListOK_Error(t *testing.T) {
	resp := NewHaGroupListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestHaGroupListOK_GetPayload(t *testing.T) {
	resp := NewHaGroupListOK()
	assert.Nil(t, resp.GetPayload())
}
