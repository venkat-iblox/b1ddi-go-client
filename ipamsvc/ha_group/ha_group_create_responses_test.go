package ha_group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHaGroupCreateCreated_Error(t *testing.T) {
	resp := NewHaGroupCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestHaGroupCreateCreated_GetPayload(t *testing.T) {
	resp := NewHaGroupCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
