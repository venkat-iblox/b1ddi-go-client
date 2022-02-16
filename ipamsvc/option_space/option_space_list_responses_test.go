package option_space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionSpaceListOK_Error(t *testing.T) {
	resp := NewOptionSpaceListOK()
	assert.NotEmpty(t, resp.Error())
}

func TestOptionSpaceListOK_GetPayload(t *testing.T) {
	resp := NewOptionSpaceListOK()
	assert.Nil(t, resp.GetPayload())
}
