package option_group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionGroupUpdateCreated_Error(t *testing.T) {
	resp := NewOptionGroupUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionGroupUpdateCreated_GetPayload(t *testing.T) {
	resp := NewOptionGroupUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
