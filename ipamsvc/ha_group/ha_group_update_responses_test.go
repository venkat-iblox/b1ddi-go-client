package ha_group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHaGroupUpdateCreated_Error(t *testing.T) {
	resp := NewHaGroupUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestHaGroupUpdateCreated_GetPayload(t *testing.T) {
	resp := NewHaGroupUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
