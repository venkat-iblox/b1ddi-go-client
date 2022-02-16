package option_filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionFilterUpdateCreated_Error(t *testing.T) {
	resp := NewOptionFilterUpdateCreated()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionFilterUpdateCreated_GetPayload(t *testing.T) {
	resp := NewOptionFilterUpdateCreated()
	assert.Nil(t, resp.GetPayload())
}
