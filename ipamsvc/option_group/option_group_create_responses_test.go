package option_group

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionGroupCreateCreated_Error(t *testing.T) {
	resp := NewOptionGroupCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionGroupCreateCreated_GetPayload(t *testing.T) {
	resp := NewOptionGroupCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
