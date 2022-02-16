package option_filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionFilterCreateCreated_Error(t *testing.T) {
	resp := NewOptionFilterCreateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionFilterCreateCreated_GetPayload(t *testing.T) {
	resp := NewOptionFilterCreateCreated()
	assert.Nil(t, resp.GetPayload())
}
