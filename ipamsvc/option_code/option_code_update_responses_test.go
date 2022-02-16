package option_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionCodeUpdateCreated_Error(t *testing.T) {
	resp := NewOptionCodeUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionCodeUpdateCreated_GetPayload(t *testing.T) {
	resp := NewOptionCodeUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
