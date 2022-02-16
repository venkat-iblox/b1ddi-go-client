package option_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionCodeCreateCreated_Error(t *testing.T) {
	resp := NewOptionCodeCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionCodeCreateCreated_GetPayload(t *testing.T) {
	resp := NewOptionCodeCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
