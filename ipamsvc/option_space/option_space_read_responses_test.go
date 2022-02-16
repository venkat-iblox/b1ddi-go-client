package option_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionSpaceReadOK_Error(t *testing.T) {
	resp := NewOptionSpaceReadOK()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionSpaceReadOK_GetPayload(t *testing.T) {
	resp := NewOptionSpaceReadOK()
	assert.Nil(t, resp.GetPayload())
}
